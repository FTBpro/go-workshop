package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
	"github.com/FTBpro/go-workshop/coolfacts/inmem"
)

func main() {
	fmt.Println("Hello, Server!")

	factsRepo := inmem.NewFactsRepository()
	service := coolfact.NewService(factsRepo)
	server := NewServer(service)

	log.Println("starting server on port 9002")
	log.Println("You can go to http://127.0.0.1:9002/ping")
	if err := http.ListenAndServe(":9002", server); err != nil {
		panic(fmt.Errorf("server crashed! err: %w", err))
	}
}
