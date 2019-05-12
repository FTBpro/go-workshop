package main

import (
	"log"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/exercise6/fact"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/exercise6/http"
	"github.com/FTBpro/go-workshop/coolfacts/exercise6/mentalfloss"
)

func main() {
	factsStore := fact.Store{}
	handlerer := facthttp.FactsHandler{
		FactStore: &factsStore,
	}

	mf := mentalfloss.Mentalfloss{}
	err := updateFacts(mf, &factsStore)

	http.HandleFunc("/ping", handlerer.Ping)
	http.HandleFunc("/facts", handlerer.Facts)

	log.Println("starting server")
	err = http.ListenAndServe(":9002", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func updateFacts(mf mentalfloss.Mentalfloss, factsStore *fact.Store) error {
	facts, err := mf.Facts()
	if err != nil {
		log.Fatal("can't reach mentalfloss: ", err)
	}
	for _, f := range facts {
		factsStore.Add(f)
	}
	return err
}