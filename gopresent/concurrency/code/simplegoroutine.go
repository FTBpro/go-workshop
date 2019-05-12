package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello Gophers!")
	for i := 1; i <= 5; i++ {
		func(j int) {
			fmt.Printf("Hello from Gopher #%d\n", j)
			time.Sleep(time.Second)
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println("finished")
}
