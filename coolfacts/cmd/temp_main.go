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
	lf *facts.ListFacts
	fc *facts.FactCreator
}

func main() {
	writeError := func(w http.ResponseWriter) {
		b, _ := json.Marshal("ERROR")
		w.Write(b)
	}

	s := facts.NewStore()
	p := facts.NewParser()
	r := facts.NewRetriever(s, p)
	lf := facts.NewListrFacts(writeError, s)
	fc := facts.NewFactCreator(writeError, p, s)
	handlerer := Handlerer{lf, fc}

	tk := time.NewTicker(time.Second * 5)
	ctx, closer := context.WithCancel(context.Background())
	defer closer()
	go func(c context.Context) {
		for {
			select {
			case <-tk.C:
				if err := r.RetrieveFacts(); err != nil {
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
		h.fc.PostFactHandler(w, r)
	case "GET":
		h.lf.PollFactHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (h *Handlerer) PingHandler(w http.ResponseWriter, r *http.Request) {
	//r.URL.Query().Get("key")
	b, _ := json.Marshal("PONG")
	w.Write(b)
}
