package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type router struct {
	router *httprouter.Router
}

func NewRouter() *httprouter.Router {
	return httprouter.New()
}

func (r *router) Handle(method, path string, handler http.Handler) {
	r.router.Handler(method, path, handler)
}
