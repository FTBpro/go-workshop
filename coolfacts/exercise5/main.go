package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

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

func main() {
	factsStore := store{}
	mf := mentalfloss{}
	facts, err := mf.Facts()
	if err != nil {
		log.Fatal("can't reach mentalfloss: ", err)
	}
	for _, f := range facts {
		factsStore.add(f)
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
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
	})

	http.HandleFunc("/facts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Add("Content-Type", "text/html")

			tmpl, err := template.New("facts").Parse(newsTemplate)
			if err != nil {
				errMessage := fmt.Sprintf("error ghttp template writing: %v", err)
				http.Error(w, errMessage, http.StatusInternalServerError)
				return
			}
			tmpl.Execute(w, factsStore.getAll())
			return
		}
		if r.Method == http.MethodPost {
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

			f := fact{
				Image:       req.Image,
				Url:         req.Url,
				Description: req.Description,
			}

			factsStore.add(f)
			w.Write([]byte("SUCCESS"))
		}
	})

	log.Println("starting server")
	err = http.ListenAndServe(":9002", nil)
	if err != nil {
		log.Fatal(err)
	}
}
