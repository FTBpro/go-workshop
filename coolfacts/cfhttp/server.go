package cfhttp

import (
	"fmt"
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

func NewServer(factsService FactsService) *server {
	// TODO: returns an initializes server with the factsService
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
	// TODO
	// 1. write status header 200 using constant http.StatusOK
	// 3. write ping
}

func (s *server) HandleGetFacts(w http.ResponseWriter) {
	facts, err := s.factsService.GetFacts()
	if err != nil {
		s.HandleError(w, fmt.Errorf("server.GetFactsHandler: %w", err))
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

func (s *server) HandleNotFound(w http.ResponseWriter, err error) {
	// TODO:
	// 1. write status header 404 using
	// 2. set content type application/json
	// 3. write json indicating an error:
	//   {
	//       error: <the error message>
	//   }
}

func (s *server) HandleError(w http.ResponseWriter, err error) {
	// TODO:
	// 1. write status header 500
	// 2. set content type application/json
	// 3. write json indicating an error:
	//   {
	//       error: <the error message>
	//   }
}
