package main

import (
	"fmt"
	"math"
)

// start OMIT

func pow(x float64, n float64, lim float64) {
	if v := math.Pow(x, n); v < lim {
		fmt.Printf("%.f ^ %.f = %.f", x, n, v)
	}
}

func main() {
	pow(3, 2, 10)
}

// end OMIT
