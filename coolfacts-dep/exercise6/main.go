package main

import (
	"log"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/exercise6/fact"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/exercise6/http"
	"github.com/FTBpro/go-workshop/coolfacts/exercise6/mentalfloss"
)

func main() {
	factsRepo := fact.Repository{}
	handlerer := facthttp.FactsHandler{
		FactRepo: &factsRepo,
	}

	mentalflossProvider := mentalfloss.Mentalfloss{}
	err := updateFacts(mentalflossProvider, &factsRepo)

	http.HandleFunc("/ping", handlerer.Ping)
	http.HandleFunc("/facts", handlerer.Facts)

	log.Println("starting server")
	err = http.ListenAndServe(":9002", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func updateFacts(mf mentalfloss.Mentalfloss, factsRepo *fact.Repository) error {
	facts, err := mf.Facts()
	if err != nil {
		log.Fatal("can't reach mentalfloss: ", err)
	}
	for _, f := range facts {
		factsRepo.Add(f)
	}
	return err
}
