package main

import "fmt"

// start OMIT
func add(x int, y int) int {
	return x + y
}

func mult(x int, y int) int {
	return x * y
}

func doSomething(x int, y int, action func(int, int) int) {
	fmt.Printf("x = %d, y = %d, action(x, y) = %d\n", x, y, action(x, y))
}

func main() {
	doSomething(3, 4, add)
	doSomething(3, 4, mult)
}

// end OMIT
