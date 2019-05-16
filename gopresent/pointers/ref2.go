package main

import (
	"fmt"
)

type TVShow struct {
	Name    string
	Seasons int
}

// start OMIT
func (s *TVShow) BumpSeasons() { // HL
	fmt.Printf("value of s before:\t%v\n", *s)

	s.Seasons += 1

	fmt.Printf("value of s after:\t%v\n", *s)
	fmt.Printf("address of s:\t%p\n", s)
}

func main() {
	a := TVShow{Name: "a", Seasons: 42}

	fmt.Printf("value of a before:\t%v\n", a)

	a.BumpSeasons()

	fmt.Printf("value of a after:\t%v\n", a)
	fmt.Printf("address of a:\t%p\n", &a)
}

// end OMIT
