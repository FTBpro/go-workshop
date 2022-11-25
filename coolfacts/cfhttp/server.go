package cfhttp

import (
	"fmt"
	"net/http"
	"strings"
)

// TODO: add methods declerations
// 1. getFacts - returns a slice of fact.Fact and an error
type Service interface {
}

type server struct {
	// TODO: add service field
}

// TODO: returns an initializes server
func NewServer(service Service) *server {
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		switch strings.ToLower(r.URL.Path) {
		case "/ping":
			s.PingHandler(w, r)
		case "/facts":
			s.GetFactsHandler(w, r)
		}
	}
}

func (s *server) PingHandler(w http.ResponseWriter, r *http.Request) {
	// TODO
	// 1. write status header 200 using constant http.StatusOK
	// 3. write ping
}

func (s *server) GetFactsHandler(w http.ResponseWriter, r *http.Request) {
	facts, err := s.service.GetFacts()
	if err != nil {
		s.ErrorHandler(w, fmt.Errorf("server.GetFactsHandler: %w", err))
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

func (s *server) NotFoundHandler(w http.ResponseWriter, err error) {
	// TODO:
	// 1. write status header 404 using
	// 2. set content type application/json
	// 3. write json indicating an error:
	//   {
	//       error: <the error message>
	//   }
}

func (s *server) ErrorHandler(w http.ResponseWriter, err error) {
	// TODO:
	// 1. write status header 500
	// 2. set content type application/json
	// 3. write json indicating an error:
	//   {
	//       error: <the error message>
	//   }
}
