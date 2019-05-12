package main

import (
	"fmt"
)

// start OMIT
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

func main() {
	fmt.Printf("%+v", store)
}

// end OMIT
