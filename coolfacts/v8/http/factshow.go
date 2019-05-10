package http

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/FTBpro/go-workshop/coolfacts/v8/fact"
)

type WriteError func(w http.ResponseWriter)

type Parser interface {
	ParseFromPolling(b []byte) ([]fact.Fact, error)
}

var newsTemplate = `
<html>
	<head>
		<title>Coolfacts</title>
	</head>
	<link rel="stylesheet" href="https://github.com/FTBpro/go-workshop/blob/master/coolfacts/styles.css">
<body>
	<h1>Did you know?</h1>
	<article>
		<a href="http://mentalfloss.com/api{{.Url}}">
				<h3>{{.Description}}</h3>
				<img src="{{.Image}}" width="100%" />
		</a>
	</article>
</body>
</html>`

func FactShowHandler(store factStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "no http handler found", http.StatusNotFound)
			return
		}
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
