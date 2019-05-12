package main

import "fmt"

// start OMIT

func main() {
	names := [4]string{"John", "Paul", "George", "Rob"}
	fmt.Println(names)

	namesCopy := names
	names[1] = "XXX"

	fmt.Println(names)
	fmt.Println(namesCopy)
}

// end OMIT
