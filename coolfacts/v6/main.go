package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/v6/fact"
	"github.com/FTBpro/go-workshop/coolfacts/v6/mentalfloss"
)

func main() {
	factsStore := fact.Store{}
	mf := mentalfloss.Mentalfloss{}
	handlerer := &Handlerer{&factsStore}

	err := updateFacts(mf, &factsStore)

	http.HandleFunc("/ping", handlerer.Ping)
	http.HandleFunc("/facts", handlerer.Facts)

	fmt.Println("starting server")
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