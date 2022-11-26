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
			// TODO: add create at to the response
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
	// TODO: add fields
	// this struct represent the client request body for the createFact API
	// Add fields:
	//		- image string
	//		- description string
	// Add json tags on the fields
}

func (r factRequest) ToCoolFact() coolfact.Fact {
	// TODO: implement this method.
}

func (s *server) HandleCreateFact(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling createFact ...")

	// TODO:
	// 1. Read the request body into factRequest
	//		Use json.NewDecoder and Decode
	// 2. Call the service for creating a fact
	// 3. On success return status OK
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
