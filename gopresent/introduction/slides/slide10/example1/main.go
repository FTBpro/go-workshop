package main

import (
	"fmt"
	"math"
)

// start OMIT
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(-4))
}

// end OMIT
