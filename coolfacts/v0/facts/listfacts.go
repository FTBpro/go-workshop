package facts

import (
	"html/template"
	"net/http"
	"strconv"
)

type WriteError func (w http.ResponseWriter)

type Parser interface {
	ParseFromPolling(b []byte) ([]Fact, error)
}

type Store interface {
	Get(i int) Fact
	GetNext() Fact
	AppendFact(fact Fact) int
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

type ListFacts struct {
	writeError WriteError
	store Store
}

func NewListrFacts(we WriteError, s Store) *ListFacts {
	return &ListFacts{we, s}
}

func (l ListFacts) PollFactHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.New("facts").Parse(newsTemplate)
	if err != nil {
		l.writeError(w)
		return
	}

	tmpl.Execute(w, l.fact(r))
}

func (l *ListFacts) fact(r *http.Request) Fact{
	indexStr := r.URL.Query().Get("index")
	if indexStr != "" {
		if index, err := strconv.Atoi(indexStr); err == nil {
			return l.store.Get(index)
		}
	}
	return l.store.GetNext()
}
