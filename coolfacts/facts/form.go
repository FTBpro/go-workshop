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
        	    Fact Url:<br/><input type="text" name="url" display="block"><br/>
				Image:<br/><input type="text" name="primaryImage" display="block"><br/>
        	    Description:<br/><input type="textarea" name="fact" display="block"><br/>
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
	f.store.AppendFact(fact)
}
