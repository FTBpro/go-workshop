package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type FactsService interface {
	GetFacts(filters coolfact.Filters) ([]coolfact.Fact, error)
	CreateFact(fact coolfact.Fact) error
}

type createFactRequest struct {
	Topic       string `json:"topic"`
	Description string `json:"description"`
}

func (r createFactRequest) ToCoolFact() coolfact.Fact {
	return coolfact.Fact{
		Topic:       r.Topic,
		Description: r.Description,
		CreatedAt:   time.Now(),
	}
}

type server struct {
	factsService FactsService
}

func NewServer(factsService FactsService) *server {
	return &server{
		factsService: factsService,
	}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("incoming request", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		switch strings.ToLower(r.URL.Path) {
		case "/ping":
			s.HandlePing(w, r)
		case "/facts":
			s.HandleGetFacts(w, r)
		default:
			s.HandleNotFound(w, r)
		}
	case http.MethodPost:
		switch strings.ToLower(r.URL.Path) {
		case "/facts":
			s.HandleCreateFact(w, r)
		default:
			s.HandleNotFound(w, r)
		}
	default:
		s.HandleNotFound(w, r)
	}
}

func (s *server) HandlePing(w http.ResponseWriter, _ *http.Request) {
	log.Println("Handling Ping ...")

	w.WriteHeader(http.StatusOK)

	if _, err := fmt.Fprint(w, "PONG"); err != nil {
		fmt.Printf("ERROR writing to ResponseWriter: %s\n", err)
		return
	}
}

func (s *server) HandleGetFacts(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling getFact ...")

	limitString := r.URL.Query().Get("limit")
	if limitString == "" || limitString == "0" {
		err := fmt.Errorf("limit is mandatory int")
		s.HandleBadRequest(w, err)
	}

	limit, err := strconv.Atoi(limitString)
	if err != nil {
		err = fmt.Errorf("HandleGetFacts limit not int")
		s.HandleBadRequest(w, err)
	}

	filters := coolfact.Filters{
		Topic: r.URL.Query().Get("topic"),
		Limit: limit,
	}

	facts, err := s.factsService.GetFacts(filters)
	if err != nil {
		s.HandleError(w, fmt.Errorf("server.HandleGetFacts: %w", err))
		return
	}

	// we first format the facts to map[string]interface.
	formattedFacts := make([]map[string]interface{}, len(facts))
	for i, coolFact := range facts {
		formattedFacts[i] = map[string]interface{}{
			"topic":       coolFact.Topic,
			"description": coolFact.Description,
			"createdAt":   coolFact.CreatedAt,
		}
	}

	response := map[string]interface{}{
		"facts": formattedFacts,
	}

	// write status and content-type
	// status must be written before the body
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	// write the body. We use json encoding
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("HandleGetFacts ERROR writing response: %s", err)
	}
}

func (s *server) HandleCreateFact(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling createFact ...")

	var request createFactRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		err = fmt.Errorf("server.HandleCreateFact failed to decode request: %s", err)
		s.HandleError(w, err)
		return
	}

	if err := s.factsService.CreateFact(request.ToCoolFact()); err != nil {
		err = fmt.Errorf("server.HandleCreateFact: %s", err)
		s.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *server) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling notFound ...")

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"error": fmt.Sprintf("path %s %s not found", r.Method, r.URL.Path),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		err = fmt.Errorf("HandleNotFound failed to decode: %s", err)
		s.HandleError(w, err)
	}
}

func (s *server) HandleError(w http.ResponseWriter, err error) {
	log.Println("Handling error ...")

	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"error": err.Error(),
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("HandleGetFacts ERROR writing response: %s", err)
	}
}

func (s *server) HandleBadRequest(w http.ResponseWriter, err error) {
	log.Println("Handling Bad Request ...")

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"error": err.Error(),
	}
	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("HandleBadRequest: %s", err)
	}
}
