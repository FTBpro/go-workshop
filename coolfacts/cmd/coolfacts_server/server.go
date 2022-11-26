package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type FactsService interface {
	GetFacts() ([]coolfact.Fact, error)
	CreateFact(fact coolfact.Fact) error
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
			s.HandlePing(w)
		case "/facts":
			s.HandleGetFacts(w)
		default:
			err := fmt.Errorf("path %q wasn't found", r.URL.Path)
			s.HandleNotFound(w, err)
		}
	case http.MethodPost:
		switch "/facts" {
		case "/facts":
			s.HandleCreateFact(w, r)
		default:
			err := fmt.Errorf("path %q wasn't found", r.URL.Path)
			s.HandleNotFound(w, err)
		}
	default:
		err := fmt.Errorf("method %q is not allowed", r.Method)
		s.HandleNotFound(w, err)
	}
}

func (s *server) HandlePing(w http.ResponseWriter) {
	log.Println("Handling Ping ...")

	w.WriteHeader(http.StatusOK)

	if _, err := fmt.Fprint(w, "PONG"); err != nil {
		fmt.Printf("ERROR writing to ResponseWriter: %s\n", err)
		return
	}
}

func (s *server) HandleGetFacts(w http.ResponseWriter) {
	log.Println("Handling getFact ...")

	facts, err := s.factsService.GetFacts()
	if err != nil {
		s.HandleError(w, fmt.Errorf("server.GetFactsHandler: %w", err))
	}

	// we first format the facts to map[string]interface.
	formattedFacts := make([]map[string]interface{}, len(facts))
	for i, coolFact := range facts {
		formattedFacts[i] = map[string]interface{}{
			"image":       coolFact.Image,
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

type factRequest struct {
	Image       string `json:"image"`
	Description string `json:"description"`
}

func (r factRequest) ToCoolFact() coolfact.Fact {
	return coolfact.Fact{
		Image:       r.Image,
		Description: r.Description,
	}
}

func (s *server) HandleCreateFact(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling createFact ...")

	var request factRequest
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

func (s *server) HandleNotFound(w http.ResponseWriter, err error) {
	log.Println("Handling notFound ...")

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]string{
		"error": err.Error(),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("HandleGetFacts ERROR writing response: %s", err)
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
