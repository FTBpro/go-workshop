package main

import (
	"encoding/json"
	"io/ioutil"
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

var store Store = Store{
	TVShows: []TVShow{
		TVShow{
			Name:    "Game of Thrones",
			Seasons: 8,
		},
	},
}

// start OMIT
func postTVShowsHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	tvShow := TVShow{}
	if err := json.Unmarshal(bodyBytes, &tvShow); err != nil {
		log.Fatal(err)
	}
	store.TVShows = append(store.TVShows, tvShow)

	w.Write([]byte("Successfully added to store"))
}

func main() {
	http.HandleFunc("/tvshows", postTVShowsHandler)
	log.Println("Listen and Serve on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// end OMIT
