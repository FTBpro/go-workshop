package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

var newsTemplate = `<!DOCTYPE html>
<html>
	<head>
		<title>Coolfacts</title>	
	<style>
/* copied from coolfacts/style.css */ 
body {
    font-family: Helvetica, Arial, sans-serif;
    color: #26323d;
    max-width: 720px;
    margin: auto;
}

article {
    border: 1px solid #0095c4;
    border-radius: 4px;
    max-width: 320px;
    text-align: center;
}

article h3 {
    font-weight: normal;
}

article a {
    color: #26323d;
}

article a:hover {
    color: #f16957;
}

article img {
    border-radius: 4px;
}
	</style>
</head>

<body>
	<h1>Hear These Amazing Facts!</h1>
	{{ range . }}
	<article>
			<h3>{{.Description}}</h3>
			<img src="{{.Image}}" width="100%" />
	</article>
	{{ end }}
</body>

</html>`

func main() {
	factsStore := store{}
	mf := mentalfloss{}
	facts, err := mf.Facts()
	if err != nil {
		log.Fatal("can't reach mentalfloss: ", err)
	}
	for _, f := range facts {
		factsStore.add(f)
	}

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "no http handler found", http.StatusNotFound)
			return
		}
		w.Header().Add("Content-Type", "text/plain")
		_, err := fmt.Fprint(w, "PONG")
		if err != nil {
			errMessage := fmt.Sprintf("error writing response: %v", err)
			http.Error(w, errMessage, http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/facts", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			w.Header().Add("Content-Type", "text/html")

			tmpl, err := template.New("facts").Parse(newsTemplate)
			if err != nil {
				errMessage := fmt.Sprintf("error ghttp template writing: %v", err)
				http.Error(w, errMessage, http.StatusInternalServerError)
				return
			}

			facts := factsStore.getAll()
			err = tmpl.Execute(w, facts)
			if err != nil {
				errMessage := fmt.Sprintf("failed parse template: %v", err)
				http.Error(w, errMessage, http.StatusInternalServerError)
				return
			}

			return
		}
		if r.Method == http.MethodPost {
			b, err := ioutil.ReadAll(r.Body)
			if err != nil {
				errMessage := fmt.Sprintf("error read from body: %v", err)
				http.Error(w, errMessage, http.StatusInternalServerError)
				return
			}
			var req struct {
				Image       string `json:"image"`
				Description string `json:"description"`
			}
			err = json.Unmarshal(b, &req)
			if err != nil {
				errMessage := fmt.Sprintf("error parsing fact: %v", err)
				http.Error(w, errMessage, http.StatusBadRequest)
			}

			f := fact{
				Image:       req.Image,
				Description: req.Description,
			}

			factsStore.add(f)
			w.Write([]byte("SUCCESS"))
		}
	})

	log.Println("starting server")
	err = http.ListenAndServe(":9002", nil)
	if err != nil {
		log.Fatal(err)
	}
}
