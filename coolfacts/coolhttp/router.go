package coolhttp

import (
	"net/http"
)

type router struct {
	notFoundHandler http.HandlerFunc
}

func NewRouter() *router {
	return &router{}
}

func (r *router) SetNotFoundHandler(handler http.HandlerFunc) {
	r.notFoundHandler = handler
}

func (r *router) Handle(method, path string, handler http.HandlerFunc) {
	// TODO: implement
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// TODO: implement
}
