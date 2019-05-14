package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TVShow struct {
	Name    string
	Seasons int
}

type Store struct {
	TVShows []TVShow
}

// start OMIT
var store Store = Store{
	TVShows: []TVShow{
		TVShow{
			Name:    "Game of Thrones",
			Seasons: 8,
		},
	},
}

func getTVShowsHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(store.TVShows)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(j)
}

func main() {
	http.HandleFunc("/tvshows", getTVShowsHandler)
	log.Println("Listen and Serve on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// end OMIT
