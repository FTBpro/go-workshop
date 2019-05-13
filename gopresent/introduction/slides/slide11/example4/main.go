package main

import "fmt"

// start OMIT

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

// end OMIT
