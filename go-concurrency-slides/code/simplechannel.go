package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan string)
	wg := new(sync.WaitGroup)
	fmt.Println("main: first print")

	wg.Add(1)
	go func(cs chan string) {
		fmt.Println("goroutine: first print")
		msg := <-cs //Blocking until message is received
		fmt.Printf("goroutine: second print - received message: %v\n", msg)
		wg.Done()
	}(ch)

	fmt.Println("main: second print")
	//Send a message
	ch <- "tester"

	wg.Wait()
	//Close the channel
	close(ch)
	fmt.Println("main: last print")
}
