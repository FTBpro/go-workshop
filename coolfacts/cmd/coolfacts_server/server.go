package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type FactsService interface {
	GetFacts() ([]coolfact.Fact, error)
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
	default:
		err := fmt.Errorf("method %q is not allowed", r.Method)
		s.HandleNotFound(w, err)
	}
}

func (s *server) HandlePing(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)

	_, err := fmt.Fprint(w, "PONG")
	if err != nil {
		fmt.Printf("ERROR writing to ResponseWriter: %s\n", err)
		return
	}
}

func (s *server) HandleGetFacts(w http.ResponseWriter) {
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

func (s *server) HandleNotFound(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"error": err,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("HandleGetFacts ERROR writing response: %s", err)
	}
}

func (s *server) HandleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"error": err,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		fmt.Printf("HandleGetFacts ERROR writing response: %s", err)
	}
}
