package main

import (
	"fmt"
)

// start OMIT
func bigNum(x int) {
	if x > 1000 {
		fmt.Printf("%d - thats a big number!", x)
	}
}

func main() {
	bigNum(1001)
}

// end OMIT
