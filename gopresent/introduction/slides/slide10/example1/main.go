package main

import (
	"fmt"
	"math"
)

// start OMIT
func sqrti(x float64) {
	if x < 0 {
		fmt.Printf("%.fi", math.Sqrt(-x))
	}
}

func main() {
	sqrti(-4)
}

// end OMIT
