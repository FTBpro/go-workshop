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
	factsStore := inmem.FactStore{}
	handlerer := facthttp.FactsHandler{
		FactStore: &factsStore,
	}

	mf := mentalfloss.Mentalfloss{}
	service := fact.NewService(&factsStore, &mf)

	if err := service.UpdateFacts(); err != nil {
		log.Fatal("couldn't update facts, try later", err.Error())
	}
	ctx, closer := context.WithCancel(context.Background())
	defer closer()
	updateFactsWithTicker(ctx, service.UpdateFacts)

	http.HandleFunc("/ping", handlerer.Ping)
	http.HandleFunc("/facts", handlerer.Facts)

	log.Println("start server")
	log.Fatal(http.ListenAndServe(":9002", nil))
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
