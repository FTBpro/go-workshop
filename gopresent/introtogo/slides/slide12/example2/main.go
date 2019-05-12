package main

import "fmt"

// start OMIT

func main() {
	names := [4]string{"John", "Paul", "George"}

	a := names[0:2]
	fmt.Println(a)

	a[0] = "XXX"
	fmt.Println(a)

	fmt.Println(names)
}

// end OMIT
