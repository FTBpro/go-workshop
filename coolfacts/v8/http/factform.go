package http

import (
	"html/template"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/v8/fact"
)

var formTemplate = `
<html>
	<head>
		<title>Coolfacts</title>
	</head>
	<style>
body {
	font-family: Helvetica, Arial, sans-serif;
	color: #26323d;
  max-width: 720px;
  margin: auto;
}

article {
	border: 1px solid #0095c4;
	border-radius: 4px;
	max-width: 256px;
	text-align: center;
}

a {
	color: #26323d;
}
a:hover {
	color: #f16957;
}
img {
	border-radius: 4px;
}
	</style>
	<body>
	{{if .Success}}
		<h1>Fact created! to show fact <a href="http://localhost:9002/facts?index={{.Index}}">click here</a></h1>
	{{else}}
				<form action="/facts/new" method="post">
						<label for="url">Fact Url:</label>
						<input id="url" type="text" name="url" display="block"><br/>
						<label for="image">Image:</label>
            <input id="image" type="text" name="primaryImage" display="block"><br/>
						<label for="fact">Description:</label>
						<input id="fact" type="textarea" name="fact" display="block"><br/>
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
		}
		switch r.Method {
		case http.MethodGet:
			if err = templ.Execute(w, nil); err != nil {
				writeError(w, err)
			}
		case http.MethodPost:
			url := r.FormValue("url")
			description := r.FormValue("fact")
			image := r.FormValue("primaryImage")

			index := store.Append(fact.Fact{Image: image, Url: url, Description: description})

			err := templ.Execute(w, struct {
				Success bool
				Index   int
			}{true, index})
			if err != nil {
				writeError(w, err)
			}
		default:
			http.Error(w, "no http handler found", http.StatusNotFound)
		}

	}
}
