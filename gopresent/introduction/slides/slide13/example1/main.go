package main

import "fmt"

// start OMIT

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2^%d = %d\n", i, v)
	}
}

// end OMIT
