package facts

import (
	"fmt"
	"html/template"
	"net/http"
)

type WriteError func (w http.ResponseWriter)

type Parser interface {
	ParseFromPolling(b []byte) ([]Fact, error)
	ParseFromCreate(b []byte) (Fact, error)
}

type Store interface {
	Get() []Fact
	Set(data []Fact)
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
	fact, err := l.getFact()
	if err != nil {
		l.writeError(w)
		return
	}

	tmpl, err := template.New("facts").Parse(newsTemplate)
	if err != nil {
		l.writeError(w)
		return
	}

	tmpl.Execute(w, fact)
}

func (l *ListFacts) getFact() (Fact, error) {
	if len(l.store.Get()) > 0 {
		return l.store.Get()[0], nil
	}
	return Fact{}, fmt.Errorf("cache empty")
}
