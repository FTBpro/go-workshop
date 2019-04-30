package http

import (
	"github.com/FTBpro/go-workshop/coolfacts/fact"
	"html/template"
	"net/http"
)

var formTemplate = `<html>
    <head>
    <title></title>
    </head>
    <body>
		{{if .Success}}
			<h1>Fact created! index = {{.Index}}</h1>
		{{else}}
        	<form action="/facts/new" method="post">
        	    Fact Url:<br/><input type="text" name="url" display="block"><br/>
				Image:<br/><input type="text" name="primaryImage" display="block"><br/>
        	    Description:<br/><input type="textarea" name="fact" display="block"><br/>
        	    <input type="submit" value="Create">
        	</form>
		{{end}}
    </body>
</html>`

type factStore interface {
	Get(i int) fact.Fact
	GetNext() fact.Fact
	Append(fact fact.Fact) int
}

func FactFormHandler(store factStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		templ, err := template.New("formFact").Parse(formTemplate)
		if err != nil {
			writeError(w, err)
			return
		}

		if r.Method != http.MethodPost {
			if err = templ.Execute(w, nil); err != nil {
				writeError(w, err)
			}
			return
		}

		url := r.FormValue("url")
		description := r.FormValue("fact")
		image := r.FormValue("primaryImage")

		index := store.Append(fact.Fact{image, url, description})

		err = templ.Execute(w, struct {
			Success bool
			Index   int
		}{true, index})
		if err != nil {
			writeError(w, err)
		}
	}
}

// private

func writeError(w http.ResponseWriter, err error) {
	http.Error(w, err.Error(), http.StatusInternalServerError)
}