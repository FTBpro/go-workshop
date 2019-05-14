package main

import (
	"fmt"
	"math"
)

// start OMIT

func bigPow(x float64, n float64) {
	if v := math.Pow(x, n); v > 1000 {
		fmt.Printf("%.f ^ %.f = %.f thats a big pow!", x, n, v)
	}
}

func main() {
	bigPow(2, 10)
}

// end OMIT
