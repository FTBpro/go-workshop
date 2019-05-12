package main

import "fmt"

var m = map[string]int{
	"Game of Thrones": 8,
	"The Simpsons":    30,
}

// start OMIT

func main() {
	m["Brooklyn 99"] = 6
	fmt.Printf("The show Brooklyn 99 has %d seasons\n", m["Brooklyn 99"])

	fmt.Printf("The show Breaking bad has %d seasons\n", m["Breaking bad"])

	v, ok := m["Game of Thrones"]
	fmt.Printf("The value %v present? %v\n", v, ok)
}

// end OMIT
