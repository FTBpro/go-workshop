package main

import  (
	"context"
	"fmt"
	"github.com/FTBpro/go-workshop/coolfacts/fact"
	facthttp "github.com/FTBpro/go-workshop/coolfacts/http"
	"github.com/FTBpro/go-workshop/coolfacts/inmem"
	"github.com/FTBpro/go-workshop/coolfacts/mentalfloss"
	"log"
	"net/http"
	"time"
)

func main() {
	store := inmem.NewFactStore()
	retriever := mentalfloss.NewRetriever()
	service := fact.NewService(store, retriever)

	if err := service.UpdateFacts(); err != nil {
		log.Fatal("couldn't update facts, try later", err.Error())
	}

	startTickerFactsUpdate(service.UpdateFacts)

	http.HandleFunc("/ping", PingHandler)
	http.HandleFunc("/facts", facthttp.FactShowHandler(store))
	http.HandleFunc("/facts/new", facthttp.FactFormHandler(store))

	log.Println("start server")
	log.Fatal(http.ListenAndServe(":9002", nil))
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	if _, err := fmt.Fprintf(w, "PONG"); err != nil {
		msg := fmt.Sprintf("coulfn't ping? %v", err)
		http.Error(w, msg, http.StatusInternalServerError)
	}
}

// private

func startTickerFactsUpdate(updateFunc func() error) {
	tk := time.NewTicker(time.Minute * 5)
	ctx, closer := context.WithCancel(context.Background())
	defer closer()
	go func(c context.Context) {
		for {
			select {
			case <-tk.C:
				if err := updateFunc(); err != nil {
					log.Printf("Error = %v", err)
				}
			case <-c.Done():
				return
			}

		}
	}(ctx)
}