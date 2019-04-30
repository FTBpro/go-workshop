package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FTBpro/go-workshop/coolfacts/facts"
	"log"
	"net/http"
	"time"
)

type Handlerer struct {
	listFacts   *facts.ListFacts
	factCreator *facts.FactCreator
}

func main() {
	writeError := func(w http.ResponseWriter) {
		b, _ := json.Marshal("ERROR")
		w.Write(b)
	}

	store := facts.NewStore()
	parser := facts.NewParser()
	retriever := facts.NewRetriever(store, parser)
	listFacts := facts.NewListrFacts(writeError, store)
	factCreator := facts.NewFactCreator(writeError, parser, store)
	handlerer := Handlerer{listFacts, factCreator}

	tk := time.NewTicker(time.Second * 5)
	ctx, closer := context.WithCancel(context.Background())
	defer closer()
	go func(c context.Context) {
		for {
			select {
			case <-tk.C:
				if err := retriever.RetrieveFacts(); err != nil {
					fmt.Printf("Error = %v", err)
				}
			case <-c.Done():
				return
			}

		}
	}(ctx)

	http.HandleFunc("/", handlerer.PingHandler)
	http.HandleFunc("/fact", handlerer.FactHandler)
	log.Fatal(http.ListenAndServe(":9002", nil))
}

func (h *Handlerer) FactHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.factCreator.PostFactHandler(w, r)
	case "GET":
		h.listFacts.PollFactHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handlerer) PingHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal("PONG")
	w.Write(b)
}
