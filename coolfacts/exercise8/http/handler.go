package http

import (
	"encoding/json"
	"fmt"
	"github.com/FTBpro/go-workshop/coolfacts/exercise8/fact"
	"html/template"
	"io/ioutil"
	"net/http"
)

type FactStore interface {
	Add(f fact.Fact)
	GetAll() []fact.Fact
}

type FactsHandler struct {
	FactStore FactStore
}

var newsTemplate = `
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

table input[type="submit"] {
    font-size: unset;
    width: 100%;
    color: #26323d;
    border: 1px solid #26323d;
    border-radius: 4px;
}

table input[type="submit"]:hover {
    color: white;
    background-color: #f16957;
    border: 1px solid #f16957;
}
	</style>
	<body>
	{{if .Success}}
		<h1>Fact created! to show fact <a href="http://localhost:9002/facts?index={{.Index}}">click here</a></h1>
	{{else}}
				<form action="/facts/new" method="post">
					<table>
						<th>Add a new fact</th>
						<tr>
							<td><label for="url">Fact Url:</label></td>
							<td><input id="url" type="text" name="url" display="block"></td>
						</tr>
						<tr>
							<td><label for="image">Image:</label></td>
            	<td><input id="image" type="text" name="primaryImage" display="block"></td>
						</tr>
						<tr>
							<td><label for="fact">Description:</label></td>
							<td><input id="fact" type="textarea" name="fact" display="block"></td>
						</tr>
						<tr><td/><td><input type="submit" value="Create"></td></tr>
					</table>
				</form>
	{{end}}
	</body>
</html>`


func (h *FactsHandler) Ping(w http.ResponseWriter, r *http.Request) {
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
}

func (h *FactsHandler) Facts(w http.ResponseWriter, r *http.Request) {
	if h.FactStore == nil {
		http.Error(w, "fact store isn't initializes", http.StatusInternalServerError)
	}

	switch r.Method {
	case http.MethodGet:
		h.showFacts(w)
		return
	case http.MethodPost:
		h.postFacts(r, w)
		return
	default:
		http.Error(w, "no http handler found", http.StatusNotFound)
	}
}

func (h *FactsHandler) postFacts(r *http.Request, w http.ResponseWriter) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		errMessage := fmt.Sprintf("error read from body: %v", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}
	var req struct {
		Image       string `json:"image"`
		Url         string `json:"url"`
		Description string `json:"description"`
	}
	err = json.Unmarshal(b, &req)
	if err != nil {
		errMessage := fmt.Sprintf("error parsing fact: %v", err)
		http.Error(w, errMessage, http.StatusBadRequest)
	}
	f := fact.Fact{
		Image:       req.Image,
		Url:         req.Url,
		Description: req.Description,
	}
	h.FactStore.Add(f)
	w.Write([]byte("SUCCESS"))
}

func (h *FactsHandler) showFacts(w http.ResponseWriter) {
	w.Header().Add("Content-Type", "text/html")
	tmpl, err := template.New("facts").Parse(newsTemplate)
	if err != nil {
		errMessage := fmt.Sprintf("error ghttp template writing: %v", err)
		http.Error(w, errMessage, http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, h.FactStore.GetAll())
}
