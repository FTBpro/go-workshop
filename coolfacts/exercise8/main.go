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
