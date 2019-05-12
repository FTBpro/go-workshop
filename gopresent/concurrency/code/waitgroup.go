package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go func(g int) {
			fmt.Printf("Hello from Gopher #%d\n", g)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println("finished")
}
