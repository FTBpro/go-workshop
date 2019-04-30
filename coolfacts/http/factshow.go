package http

import (
	"github.com/FTBpro/go-workshop/coolfacts/fact"
	"html/template"
	"net/http"
	"strconv"
)

type WriteError func(w http.ResponseWriter)

type Parser interface {
	ParseFromPolling(b []byte) ([]fact.Fact, error)
}

var newsTemplate = `<html>
                    <h1>News</h1>
                    <div>
                            <div>
                                <h3>{{.Description}}</h3>
                                <img src="{{.Image}}" width="25%" height="25%"></img>
                            </div>
                    <div>
                    </html>`


func FactShowHandler(store factStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r * http.Request) {
		tmpl, err := template.New("facts").Parse(newsTemplate)
		if err != nil {
			writeError(w, err)
			return
		}

		indexStr := r.URL.Query().Get("index")
		index, err := strconv.Atoi(indexStr)
		var f fact.Fact
		if err != nil {
			f = store.GetNext()
		} else {
			f = store.Get(index)
		}

		if err = tmpl.Execute(w, f); err != nil {
			writeError(w, err)
			return
		}
	}
}
