package main

import (
	"fmt"
)

func main(){
	a := 42
	do := func (b int) { // HL
		fmt.Printf("value of b:\t%d\n", b)
		fmt.Printf("address of b:\t%p\n", &b)
	}
	do(a) // HL

	fmt.Printf("value of a:\t%d\n", a)
	fmt.Printf("address of a:\t%p\n", &a)
}


