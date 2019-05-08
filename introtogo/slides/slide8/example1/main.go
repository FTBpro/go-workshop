package main

import "fmt"

func main() {
	// start OMIT
	type MySpcialInt int
	var m MySpcialInt = 3
	var i int = 4 // i = m -> compile error
	fmt.Printf("Type: %T Value: %v\n", m, m)
	fmt.Printf("Type: %T Value: %v\n", i, i)
	// end OMIT
}
