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
	factsStore.add(fact{
		Image:       "https://images2.minutemediacdn.com/image/upload/v1556645500/shape/cover/entertainment/D5aliXvWsAEcYoK-fe997566220c082b98030508e654948e.jpg",
		Url:         "http://example.com",
		Description: "Did you know sonic is a hedgehog?!",
	})
	factsStore.add(fact{
		Image:       "https://images2.minutemediacdn.com/image/upload/v1556641470/shape/cover/entertainment/uncropped-Screen-Shot-2019-04-30-at-122411-PM-3b804f143c543dfab4b75c81833bed1b.jpg",
		Url:         "http://example.com",
		Description: "You won't believe what happened to Arya!",
	})

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
				Image: req.Image,
				Url: req.Url,
				Description: req.Description,
			}

			factsStore.add(f)
			w.Write([]byte("SUCCESS"))
		}
	})

	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		log.Fatal(err)
	}
}
