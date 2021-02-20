package main

import (
	"fmt"
)

// start OMIT
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

func main() {
	fmt.Printf("%+v", repo)
}

// end OMIT
