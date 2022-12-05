package main

import (
	"fmt"
	"github.com/FTBpro/go-workshop/coolfacts/coolhttp"
	"log"
	"net/http"
	"time"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
	"github.com/FTBpro/go-workshop/coolfacts/inmem"
)

func main() {
	fmt.Println("Hello, Server!")

	factsRepo := inmem.NewFactsRepository(seedFacts()...)
	service := coolfact.NewService(factsRepo)
	server := NewServer(service)

	router := coolhttp.NewRouter()
	router.SetNotFoundHandler(server.HandleNotFound)

	server.RegisterRouter(router)

	log.Println("starting server on port 9002")
	log.Println("You can go to http://127.0.0.1:9002/ping")
	if err := http.ListenAndServe(":9002", router); err != nil {
		panic(fmt.Errorf("server crashed! err: %w", err))
	}
}

func seedFacts() []coolfact.Fact {
	return []coolfact.Fact{
		{
			Topic:       "Games",
			Description: "Did you know sonic is a hedgehog?!",
			CreatedAt:   time.Now(),
		},
		{
			Topic:       "TV",
			Description: "You won't believe what happened to Arya!",
			CreatedAt:   time.Now().Add(-time.Duration(1) * time.Hour),
		},
	}
}
