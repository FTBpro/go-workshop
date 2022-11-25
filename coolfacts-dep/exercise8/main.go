package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/FTBpro/go-workshop/coolfacts/exercise8/fact"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/exercise8/http"
	"github.com/FTBpro/go-workshop/coolfacts/exercise8/inmem"
	"github.com/FTBpro/go-workshop/coolfacts/exercise8/mentalfloss"
)

const (
	updateFactInterval = 5 * time.Minute
)

func main() {
	factsRepository := inmem.NewFactRepository()
	handler := facthttp.NewFactsHandler(factsRepository)

	mentalflossProvider := mentalfloss.NewProvider()
	service := fact.NewService(mentalflossProvider, factsRepository)
	if err := service.UpdateFacts(); err != nil {
		log.Fatal("couldn't update facts, try later", err.Error())
	}

	ctx, closer := context.WithCancel(context.Background())
	defer closer()
	updateFactsWithTicker(ctx, service.UpdateFacts)

	http.HandleFunc("/ping", handler.Ping)
	http.HandleFunc("/facts", handler.Facts)

	const port = ":9002"
	log.Println("started server on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// private

func updateFactsWithTicker(ctx context.Context, updateFunc func() error) {
	tk := time.NewTicker(updateFactInterval)
	go func(c context.Context) {
		for {
			select {
			case <-tk.C:
				log.Println("updating...")
				if err := updateFunc(); err != nil {
					log.Printf("error updating = %v", err)
				}
				log.Println("updated Successfully")
			case <-c.Done():
				return
			}
		}
	}(ctx)
}
