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
	// TODO: add method CreateFact
}

// TODO: add struct factRequest
// This struct should represent the client request for creating a new fact.
// The client sends JSON:
// {
//		"topic": "...",
//		"description": "..."
// }
// TODO: add method on this struct `ToCoolFact` that convert it into an entity coolfact.Fact

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

	// TODO: add case to support the create fact API
	// the expected path for creating a fact is "/paths", and the http method is POST (http.MethodPost)
	// use server method HandleCreateFact

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

func (s *server) HandleGetFacts(w http.ResponseWriter, _ *http.Request) {
	log.Println("Handling getFact ...")

	facts, err := s.factsService.GetFacts()
	if err != nil {
		s.HandleError(w, fmt.Errorf("server.GetFactsHandler: %w", err))
		return
	}

	response := s.formatGetFactsResponse(facts)

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

	// TODO:
	// 1. Read the request body into factRequest
	//		Use json.NewDecoder and Decode
	// 2. Call the service for creating a fact
	// 3. On success return status OK
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

func (s *server) formatGetFactsResponse(facts []coolfact.Fact) map[string]interface{} {
	formattedFacts := make([]map[string]interface{}, len(facts))
	for i, coolFact := range facts {
		formattedFacts[i] = map[string]interface{}{
			"topic":       coolFact.Topic,
			"description": coolFact.Description,
			// TODO: add created at to the response
		}
	}

	return map[string]interface{}{
		"facts": formattedFacts,
	}
}
