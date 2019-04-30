package facts

import (
	"html/template"
	"net/http"
)

var formTemplate = `<html>
    <head>
    <title></title>
    </head>
    <body>
		{{if .Success}}
			<h1>Fact created!</h1>
		{{else}}
        	<form action="/facts/new" method="post">
        	    FactUrl:<input type="text" name="url">
        	    Description:<input type="text" name="fact">
				Image:<input type="text" name="primaryImage">
        	    <input type="submit" value="Create">
        	</form>
		{{end}}
    </body>
</html>`

type FactForm struct {
	writeError WriteError
	store      Store
}

func NewFactForm(we WriteError, s Store) *FactForm {
	return &FactForm{we, s}
}

func (f *FactForm) FormFactHandler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.New("formFact").Parse(formTemplate)
	if err != nil {
		f.writeError(w)
	}
	if r.Method != http.MethodPost {
		templ.Execute(w, nil)
		return
	}

	url := r.FormValue("url")
	description := r.FormValue("fact")
	image := r.FormValue("primaryImage")

	f.WriteToCache(Fact{image, url, description})

	templ.Execute(w, struct{ Success bool }{true})
}

func (f *FactForm) WriteToCache(fact Fact) {
	data := f.store.Get()
	data = append([]Fact{fact}, data...)
	f.store.Set(data)
}
