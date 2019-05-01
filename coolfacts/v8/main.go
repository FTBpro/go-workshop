package main

import (
	"context"
	"github.com/FTBpro/go-workshop/coolfacts/v8/fact"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/v8/http"
	"github.com/FTBpro/go-workshop/coolfacts/v8/inmem"
	"github.com/FTBpro/go-workshop/coolfacts/v8/mentalfloss"
	"log"
	"net/http"
	"time"
)

const (
	updateFactInterval = time.Minute * 5
)

func main() {
	store := inmem.NewFactStore()
	factProvider := mentalfloss.NewProvider()
	service := fact.NewService(store, factProvider)

	if err := service.UpdateFacts(); err != nil {
		log.Fatal("couldn't update facts, try later", err.Error())
	}

	startTickerFactsUpdate(service.UpdateFacts, updateFactInterval)

	http.HandleFunc("/ping", facthttp.PingHandler)
	http.HandleFunc("/facts", facthttp.FactShowHandler(store))
	http.HandleFunc("/facts/new", facthttp.FactFormHandler(store))

	log.Println("start server")
	log.Fatal(http.ListenAndServe(":9002", nil))
}

// private

func startTickerFactsUpdate(updateFunc func() error, rate time.Duration) {
	tk := time.NewTicker(rate)
	ctx, closer := context.WithCancel(context.Background())
	defer closer()
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