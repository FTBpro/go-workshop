package cfhttp

import "net/http"

type server struct {
}

func NewServer() *server {
	return &server{}
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// TODO: if the request is ping, return PONG, otherwise return not found
}