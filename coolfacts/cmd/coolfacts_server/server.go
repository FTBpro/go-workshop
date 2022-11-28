package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type server struct{}

func NewServer() *server {
	return &server{}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("incoming request", r.Method, r.URL.Path)

	switch r.Method {
	case http.MethodGet:
		switch strings.ToLower(r.URL.Path) {
		case "/ping":
			s.HandlePing(w)
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
