package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FTBpro/go-workshop/coolfacts/v0/facts"
	"log"
	"net/http"
	"time"
)

func main() {
	writeError := func(w http.ResponseWriter) {
		b, _ := json.Marshal("ERROR")
		w.Write(b)
	}

	store := facts.NewStore()
	parser := facts.NewParser()
	retriever := facts.NewRetriever(store, parser)
	listFacts := facts.NewListrFacts(writeError, store)
	factForm := facts.NewFactForm(writeError, store)

	retriever.RetrieveFacts()
	tk := time.NewTicker(time.Minute * 5)
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

	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/facts", listFacts.PollFactHandler)
	http.HandleFunc("/facts/new", factForm.FormFactHandler)
	log.Fatal(http.ListenAndServe(":9002", nil))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	b, _ := json.Marshal("PONG")
	w.Write(b)
}
