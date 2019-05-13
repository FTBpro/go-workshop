package main

import "fmt"

// start OMIT

var m = map[string]int{
	"Game of Thrones": 8,
	"The Simpsons":    30,
}

func main() {
	for k, v := range m {
		fmt.Printf("Show name : %s, # of seasons: %d\n", k, v)
	}
}

// end OMIT
