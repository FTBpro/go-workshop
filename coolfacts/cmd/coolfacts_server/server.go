package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type server struct{}

func NewServer() *server {
	// TODO: returns an initializes server with the factsService
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

func (s *server) HandlePing(w http.ResponseWriter, _ *http.Request) {
	log.Println("Handling Ping ...")

	// TODO
	// 1. write status header 200 using constant http.StatusOK
	// 3. write ping
}

func (s *server) HandleNotFound(w http.ResponseWriter, r *http.Request) {
	log.Println("Handling notFound ...")
	// TODO:
	// 1. write status header 404 using
	// 2. set content type application/json
	// 3. write json indicating an error:
	//   {
	//       error: path <path> not not found
	//   }
}
