package main

import "fmt"

// start OMIT

func main() {
	names := []string{"John", "Paul", "George", "Rob"}
	fmt.Printf("Names slice : %v\n", names)

	names[1] = "XXX"
	fmt.Printf("Names slice : %v\n", names)

	names = append(names, "Tom")
	fmt.Printf("Names slice : %v\n", names)
}

// end OMIT
