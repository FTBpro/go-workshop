package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/FTBpro/go-workshop/coolfacts/exercise7/fact"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/exercise7/http"
	"github.com/FTBpro/go-workshop/coolfacts/exercise7/mentalfloss"
)

const (
	updateFactInterval = time.Minute * 1
)

func main() {
	factsRepo := fact.Repository{}
	handlerer := facthttp.FactsHandler{
		FactRepo: &factsRepo,
	}

	mentalflossProvider := mentalfloss.Mentalfloss{}
	updateFunc := updateFactsFunc(mentalflossProvider, &factsRepo)
	if err := updateFunc(); err != nil {
		log.Fatal(err)
	}

	ctx, closer := context.WithCancel(context.Background())
	defer closer()
	updateFactsWithTicker(ctx, updateFunc)

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

func updateFactsFunc(mf mentalfloss.Mentalfloss, factsRepo *fact.Repository) func() error {
	return func() error {
		facts, err := mf.Facts()
		if err != nil {
			log.Fatal("can't reach mentalfloss: ", err)
		}
		for _, f := range facts {
			factsRepo.Add(f)
		}
		return err
	}
}
