package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/FTBpro/go-workshop/coolfacts/v7/fact"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/v7/http"
	"github.com/FTBpro/go-workshop/coolfacts/v7/mentalfloss"
)

const (
	updateFactInterval = time.Minute * 1
)

func main() {
	factsStore := fact.Store{}
	handlerer := facthttp.FactsHandler{
		FactStore: &factsStore,
	}

	mf := mentalfloss.Mentalfloss{}
	updateFunc := updateFactsFunc(mf, &factsStore)
	if err := updateFunc(); err != nil {
		log.Fatal(err)
	}

	ctx, closer := context.WithCancel(context.Background())
	defer closer()
	updateFactsWithTicker(ctx, updateFunc)

	http.HandleFunc("/ping", handlerer.Ping)
	http.HandleFunc("/facts", handlerer.Facts)

	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":9002", nil))
}

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

func updateFactsFunc(mf mentalfloss.Mentalfloss, factsStore *fact.Store) func() error {
	return func() error {
		facts, err := mf.Facts()
		if err != nil {
			log.Fatal("can't reach mentalfloss: ", err)
		}
		for _, f := range facts {
			factsStore.Add(f)
		}
		return err
	}
}
