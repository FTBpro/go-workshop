package mentalfloss

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/v8/fact"
)

const getFactsAPI = "http://mentalfloss.com/api/facts"

type provider struct{}

func NewProvider() *provider {
	return &provider{}
}

func (r *provider) Facts() ([]fact.Fact, error) {
	resp, err := http.Get(getFactsAPI)
	if err != nil {
		return nil, fmt.Errorf("error get = %v", err)
	}

	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error readAll = %v", err)
	}

	facts, err := parseFromRawItems(b)
	if err != nil {
		return nil, fmt.Errorf("error parsing data = %v", err)
	}

	return facts, nil
}

func parseFromRawItems(b []byte) ([]fact.Fact, error) {
	items := make([]item, 0)
	if err := json.Unmarshal(b, &items); err != nil {
		return nil, err
	}

	facts := make([]fact.Fact, 0)
	for _, it := range items {
		newFact := it.ToFact()
		facts = append(facts, newFact)
	}

	return facts, nil
}
