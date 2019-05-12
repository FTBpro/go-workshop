package main

import "fmt"

var m = map[string]int{
	"Game of Thrones": 8,
	"The Simpsons":    30,
}

// start OMIT

func main() {
	m["Brooklyn 99"] = 6
	fmt.Println(m["Brooklyn 99"])

	v, ok := m["Game of Thrones"]
	fmt.Println("The value:", v, "Present?", ok)
}

// end OMIT
