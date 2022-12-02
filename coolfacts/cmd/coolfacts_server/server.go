package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type FactsService interface {
	// TODO: add methods declerations
	// 1. getFacts - returns a slice of fact.Fact and an error
}

type server struct {
	// TODO: add factsService field
}

func NewServer() *server {
	// TODO: returns an initializes server with the factsService
	return &server{}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("incoming request", r.Method, r.URL.Path)

	// TODO: add case for GET /facts, that will call to `HandleGetFacts`
	switch r.Method {
	case http.MethodGet:
		switch strings.ToLower(r.URL.Path) {
		case "/ping":
			s.HandlePing(w, r)
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

	// TODO:
	// 1. format the facts to a json response
	// 2. write status 200
	// 3. set content type application/json
	// 4. write json response:
	// 	{
	//		"facts": [
	//			{
	//				"id": "..."
	//				"description": "..."
	//			},
	//			...
	//		]
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
