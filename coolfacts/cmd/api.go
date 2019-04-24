package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type DecodeFunc func(r *http.Request) (interface{}, error)

type ProcessRequest func(interface{}) (interface{}, error)

type EncodedFunc func(interface{}) (interface{}, error)

type HttpHandler struct {
	Decoder   DecodeFunc
	Processor ProcessRequest
	Encoder   EncodedFunc
}

type RouteDefinition struct {
	Method  string
	Path    string
	Handler HttpHandler
}

type Rout interface {
	MakeRoutesDefinitions() []*RouteDefinition
}

type api struct {
}

func NewApi() *api {
	return &api{}
}

func (a *api) MakeRouts(routs []Rout) {
	for _, rout := range routs {
		routDefinitions := rout.MakeRoutesDefinitions()
		for _, routDefinition := range routDefinitions {
			http.HandleFunc(routDefinition.Path, a.MakeHandlerFunc(routDefinition))
		}
	}
	log.Fatal(http.ListenAndServe(":3901", nil))
}

func (a *api) MakeHandlerFunc(routDefinition *RouteDefinition) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		decodedData, err := routDefinition.Handler.Decoder(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		processData, err := routDefinition.Handler.Processor(decodedData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		response, err := routDefinition.Handler.Encoder(processData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		stream, err := json.Marshal(response)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Write(stream)
	}
}