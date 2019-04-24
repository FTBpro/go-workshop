package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var Cache = Store{}

type Store struct {
	Data  []MF
}

type MF struct {
	Id            string   `json:"id"`
	Url           string   `json:"url"`
	FactId        string   `json:"factId"`
	Headline      string   `json:"headline"`
	ShortHeadline string   `json:"shortHeadline"`
	Fact          string   `json:"fact"`
	FullStoryUrl  string   `json:"fullStoryUrl"`
	Tags          []string `json:"tags"`
	PrimaryImage  string   `json:"primaryImage"`
	ImageCredit   string   `json:"imageCredit"`
}

var newsTemplate = `<html>
                    <h1>News</h1>
                    <div>
                        {{range .}}
                            <div>
                                <h3>{{.Fact}}</h3>
                                <img src="{{.PrimaryImage}}" width="25%" height="25%"></img>
                            </div>
                        {{end}}
                    <div>
                    </html>`

func main() {
	tk := time.NewTicker(time.Second * 5)
	ctx, closer := context.WithCancel(context.Background())
	defer closer()
	go func(c context.Context) {
		for {
			select {
			case <-tk.C:
				if err := RetrieveFact(); err != nil {
					fmt.Printf("Error = %v", err)
				}
			case <-c.Done():
				return
			}

		}
	}(ctx)

	http.HandleFunc("/", PingHandler)
	http.HandleFunc("/fact", PollFactHandler)
	log.Fatal(http.ListenAndServe(":3902", nil))
}

func PingHandler  (w http.ResponseWriter, r *http.Request) {
	//r.URL.Query().Get("key")
	b, _ := json.Marshal("PONG")
	w.Write(b)
}

func PollFactHandler(w http.ResponseWriter, r *http.Request) {
	if len(Cache.Data) > 0 {
		tmpl, err := template.New("facts").Parse(newsTemplate)
		if err != nil {
			WriteError(w)
			return
		}
		tmpl.Execute(w, Cache.Data)
		return
	}
	WriteError(w)
}

func WriteError(w http.ResponseWriter) {
	b, _ := json.Marshal("ERROR")
	w.Write(b)
}

func getFact() []byte {
	b := make([]byte, 0)
	if len(Cache.Data) > 0 {
		b, _ = json.Marshal(Cache.Data[0])
	} else {
		b, _ = json.Marshal("ERROR")
	}
	return b
}


func RetrieveFact() error {
	resp, err := http.Get("http://mentalfloss.com/api/facts")
	if err != nil {
		return fmt.Errorf("error get = %v", err)

	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error readAll = %v", err)
	}
	data := make([]MF, 0)
	err = json.Unmarshal(b, &data)
	if err != nil {
		return fmt.Errorf("error parsing data = %v", err)
	}
	Cache.Data = data
	return nil
}