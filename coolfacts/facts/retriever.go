package facts

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type retriever struct {
	store Store
	parser Parser
}

func NewRetriever(s Store, p Parser) *retriever {
	return &retriever{s, p}
}

func (r *retriever) RetrieveFacts() error {
	resp, err := http.Get("http://mentalfloss.com/api/facts")
	if err != nil {
		return fmt.Errorf("error get = %v", err)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error readAll = %v", err)
	}

	parsedFacts, err := r.parseFacts(b)
	if err != nil {
		return fmt.Errorf("error parsing data = %v", err)
	}

	r.cacheData(parsedFacts)
	return nil
}

func (r *retriever) parseFacts(b []byte) ([]Fact, error) {
	return r.parser.ParseFromPolling(b)
}

func (r *retriever) cacheData(facts []Fact) {
	r.store.Set(facts)
}

