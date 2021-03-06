package main

import (
	"fmt"
)

type TVShow struct {
	Name    string
	Seasons int
}

// start OMIT
func (s TVShow) BumpSeasonsBUG() { // HL
	fmt.Printf("value of s before:\t%v\n", s)

	s.Seasons += 1 // HL

	fmt.Printf("value of s after:\t%v\n", s)
	fmt.Printf("address of s:\t%p\n", &s)
} // HL

func main() {
	a := TVShow{Name: "a", Seasons: 42}

	fmt.Printf("value of a before:\t%v\n", a)

	a.BumpSeasonsBUG() // HL

	fmt.Printf("value of a after:\t%v\n", a)
	fmt.Printf("address of a:\t%p\n", &a)
}

// end OMIT
