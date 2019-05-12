package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/FTBpro/go-workshop/coolfacts/v8/fact"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/v8/http"
	"github.com/FTBpro/go-workshop/coolfacts/v8/inmem"
	"github.com/FTBpro/go-workshop/coolfacts/v8/mentalfloss"
)

const (
	updateFactInterval = 5 * time.Minute
)

func main() {
	store := inmem.NewFactStore()
	factProvider := mentalfloss.NewProvider()
	service := fact.NewService(store, factProvider)

	if err := service.UpdateFacts(); err != nil {
		log.Fatal("couldn't update facts, try later", err.Error())
	}
	ctx, closer := context.WithCancel(context.Background())
	defer closer()
	startTickerFactsUpdate(ctx, service.UpdateFacts, updateFactInterval)

	http.HandleFunc("/ping", facthttp.PingHandler)
	http.HandleFunc("/facts", facthttp.FactShowHandler(store))
	http.HandleFunc("/facts/new", facthttp.FactFormHandler(store))

	const port = ":9002"
	log.Println("started server on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}

// private

func startTickerFactsUpdate(ctx context.Context, updateFunc func() error, rate time.Duration) {
	tk := time.NewTicker(rate)
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
