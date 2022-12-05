package coolhttp

import (
	"fmt"
	"net/http"
)

type methodPathToHandler map[string]http.HandlerFunc

type router struct {
	notFoundHandler http.HandlerFunc
	handles         methodPathToHandler
}

func NewRouter() *router {
	return &router{
		handles: methodPathToHandler{},
	}
}

func (r *router) SetNotFoundHandler(handler http.HandlerFunc) {
	r.notFoundHandler = handler
}

func (r *router) Handle(method, path string, handler http.HandlerFunc) {
	r.handles[fmt.Sprintf("%s|%s", method, path)] = handler
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	handler, ok := r.handles[fmt.Sprintf("%s|%s", req.Method, req.URL.Path)]
	if !ok {
		if r.notFoundHandler != nil {
			r.notFoundHandler(w, req)
		}
		return
	}

	handler(w, req)
}
