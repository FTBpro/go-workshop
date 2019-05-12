package main

import "fmt"
	
type Action interface {
	Do(int, int) int
}

func DoSomething(x, y int, action Action) {
	fmt.Printf("x = %d, y = %d, action.Do(x, y) = %d\n", x, y, action.Do(x, y))
}


// start OMIT
type adder struct{}
func (a adder) Do(x, y int) int {
	return x + y
}

type multiplier struct{}
func (m multiplier) Do(x, y int) int {
	return x * y
}

func main() {
	DoSomething(3, 4, adder{})
	DoSomething(3, 4, multiplier{})
}
// end OMIT