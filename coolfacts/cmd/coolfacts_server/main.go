package main

import (
	"fmt"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/fact"
	"github.com/FTBpro/go-workshop/coolfacts/inmem"
)

func main() {
	fmt.Println("Hello, World!")

	factsRepo := inmem.NewFactsRepository()
	service := fact.NewService(factsRepo)
	server := NewServer(service)

	fmt.Println("starting server on port 9002")
	fmt.Println("You can go to http://127.0.0.1:9002/ping")
	if err := http.ListenAndServe(":9002", server); err != nil {
		panic(fmt.Errorf("server crashed! err: %w", err))
	}
}
