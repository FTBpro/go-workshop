package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	server := http.NewServeMux()

	server.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/plain")
		_, err := fmt.Fprint(w, "PONG")
		if err != nil {
			errMessage := fmt.Sprintf("error writing response: %v", err)
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":9002", server)
	if err != nil {
		log.Fatal(err)
	}
}
