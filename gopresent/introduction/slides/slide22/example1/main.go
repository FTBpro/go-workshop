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

type Repository struct {
	TVShows []TVShow
}

var repo = Repository{
	TVShows: []TVShow{
		{
			Name:    "Game of Thrones",
			Seasons: 8,
		},
	},
}

func getTVShowsHandler(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(repo)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(j)
}

func postTVShowsHandler(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	tvShow := TVShow{}
	if err := json.Unmarshal(bodyBytes, &tvShow); err != nil {
		log.Fatal(err)
	}
	repo.TVShows = append(repo.TVShows, tvShow)

	w.Write([]byte("Successfully added to repo"))
}

// start OMIT
func tvshowsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getTVShowsHandler(w, r)
		return
	}

	if r.Method == http.MethodPost {
		postTVShowsHandler(w, r)
		return
	}

	log.Fatal("Unsupported request method")
}

func main() {
	http.HandleFunc("/tvshows", tvshowsHandler)
	log.Println("Listen and Serve on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// end OMIT
