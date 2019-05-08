package http

import (
	"encoding/json"
	"fmt"
	"github.com/FTBpro/go-workshop/coolfacts/v8/fact"
	"html/template"
	"io/ioutil"
	"net/http"
)

type FactStore interface {
	Add(f fact.Fact)
	GetAll() []fact.Fact
}

type FactsHandler struct {
	FactStore FactStore
}

var newsTemplate = `<html>
                    <h1>Facts</h1>
                    <div>
                        {{range .}}
                            <div>
                                <h3>{{.Description}}</h3>
                                <img src="{{.Image}}" width="25%" height="25%"> </img>
                            </div>
                        {{end}}
                    <div>
                    </html>`


func (h *FactsHandler) Ping(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "no http handler found", http.StatusNotFound)
		return
	}
	w.Header().Add("Content-Type", "text/plain")
	_, err := fmt.Fprint(w, "PONG")
	if err != nil {
		errMessage := fmt.Sprintf("error writing response: %v", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
	}
}

func (h *FactsHandler) Facts(w http.ResponseWriter, r *http.Request) {
	if h.FactStore == nil {
		http.Error(w, "fact store isn't initializes", http.StatusInternalServerError)
	}

	if r.Method == http.MethodGet {
		h.getFacts(w)
		return
	}
	if r.Method == http.MethodPost {
		h.postFacts(r, w)
		return
	}
	http.Error(w, "no http handler found", http.StatusNotFound)
}

func (h *FactsHandler) postFacts(r *http.Request, w http.ResponseWriter) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errMessage := fmt.Sprintf("error read from body: %v", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}
	var req struct {
		Image       string `json:"image"`
		Url         string `json:"url"`
		Description string `json:"description"`
	}
	err = json.Unmarshal(b, &req)
	if err != nil {
		errMessage := fmt.Sprintf("error parsing fact: %v", err)
		http.Error(w, errMessage, http.StatusBadRequest)
	}
	f := fact.Fact{
		Image:       req.Image,
		Url:         req.Url,
		Description: req.Description,
	}
	h.FactStore.Add(f)
	w.Write([]byte("SUCCESS"))
}

func (h *FactsHandler) getFacts(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("facts").Parse(newsTemplate)
	if err != nil {
		errMessage := fmt.Sprintf("error ghttp template writing: %v", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, h.FactStore.GetAll())
}
