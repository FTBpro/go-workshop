package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	factsStore := store{}
	factsStore.add(fact{
		Image:       "https://minutemedia-res.cloudinary.com/image/upload/v1556645500/shape/cover/entertainment/D5aliXvWsAEcYoK-fe997566220c082b98030508e654948e.jpg",
		Url:         "http://example.com",
		Description: "Did you know sonic is a hedgehog?!",
	})
	factsStore.add(fact{
		Image:       "https://minutemedia-res.cloudinary.com/image/upload/v1556641470/shape/cover/entertainment/uncropped-Screen-Shot-2019-04-30-at-122411-PM-3b804f143c543dfab4b75c81833bed1b.jpg",
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
		if r.Method != http.MethodGet {
			http.Error(w, "no http handler found", http.StatusNotFound)
			return
		}
		w.Header().Add("Content-Type", "application/json")

		b, err := json.Marshal(factsStore.getAll())
		if err != nil {
			errMessage := fmt.Sprintf("error marshaling facts : %v", err)
			http.Error(w, errMessage, http.StatusInternalServerError)
			return
		}

		_, err = w.Write(b)
		if err != nil {
			errMessage := fmt.Sprintf("error writing response: %v", err)
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		log.Fatal(err)
	}
}
