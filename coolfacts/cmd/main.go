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
	Data []MF
}

type MF struct {
	Id            string   `json:"id"`
	Url           string   `json:"url"`
	FactId        string   `json:"factId"`
	Headline      string   `json:"headline"`
	ShortHeadline string   `json:"shortHeadline"`
	FactText      string   `json:"fact"`
	FullStoryUrl  string   `json:"fullStoryUrl"`
	Tags          []string `json:"tags"`
	PrimaryImage  string   `json:"primaryImage"`
	ImageCredit   string   `json:"imageCredit"`
}

type Fact struct {
	Image       string
	Url         string
	Description string
}

var newsTemplate = `<html>
                    <h1>News</h1>
                    <div>
                        {{range .}}
                            <div>
                                <h3>{{.FactText}}</h3>
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
	http.HandleFunc("/fact", FactHandler)
	log.Fatal(http.ListenAndServe(":9002", nil))
}

func FactHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		PostFactHadnler(w, r)
	case "GET":
		PollFactHandler(w, r)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func PingHandler(w http.ResponseWriter, r *http.Request) {
	//r.URL.Query().Get("key")
	b, _ := json.Marshal("PONG")
	w.Write(b)
}

func PostFactHadnler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w)
	}
	fact := MF{}
	err = json.Unmarshal(b, &fact)
	if err != nil {
		WriteError(w)
	}
	Cache.Data = append(Cache.Data, fact)
}

func PollFactHandler(w http.ResponseWriter, r *http.Request) {
	facts, err := getFact()
	if err != nil {
		WriteError(w)
		return
	}
	tmpl, err := template.New("facts").Parse(newsTemplate)
	if err != nil {
		WriteError(w)
		return
	}
	tmpl.Execute(w, facts)
}

func WriteError(w http.ResponseWriter) {
	b, _ := json.Marshal("ERROR")
	w.Write(b)
}

func getFact() ([]MF, error) {
	if len(Cache.Data) > 0 {
		return Cache.Data, nil
	}
	return nil, fmt.Errorf("cache empty")
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

func ParseFact(mf MF) *Fact {
	return &Fact{
		Description: mf.FactText,
		Url:         mf.Url,
		Image:       mf.PrimaryImage,
	}
}
