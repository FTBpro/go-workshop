package main

import "fmt"

// start1 OMIT
type TVShow struct {
	Name    string
	Seasons int
}

// end1 OMIT

// start2 OMIT
func (s TVShow) String() {
	fmt.Printf("The TV show %s, has %d seasons", s.Name, s.Seasons)
}

// end2 OMIT

// start3 OMIT
func main() {
	got := TVShow{
		Name:    "Game of Thrones",
		Seasons: 8,
	}
	got.String()
}

// end3 OMIT
