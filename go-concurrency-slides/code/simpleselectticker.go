package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	dataChannel := make(chan int)
	ticker := time.NewTicker(time.Second)
	timeout := time.After(5 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				response, _ := http.Get("https://golang.org/")
				dataChannel <- response.StatusCode
			case <-timeout:
				ticker.Stop()
				close(dataChannel)
				return
			}
		}
	}()

	for dc := range dataChannel {
		fmt.Printf("http status is: %d\n", dc)
	}
	fmt.Println("finished")
}
