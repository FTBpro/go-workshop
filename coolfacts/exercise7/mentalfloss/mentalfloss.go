package mentalfloss

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FTBpro/go-workshop/coolfacts/exercise7/fact"
)

type Mentalfloss struct {
}

func (mf Mentalfloss) Facts() ([]fact.Fact, error) {
	resp, err := http.Get("http://mentalfloss.com/api/facts")
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error readAll: %v", err)
	}

	facts, err := parseFromRawItems(b)
	if err != nil {
		return nil, fmt.Errorf("error parsing data: %v", err)
	}

	return facts, nil
}

func parseFromRawItems(b []byte) ([]fact.Fact, error) {
	var items []struct {
		Url          string `json:"url"`
		FactText     string `json:"fact"`
		PrimaryImage string `json:"primaryImage"`
	}
	if err := json.Unmarshal(b, &items); err != nil {
		return nil, err
	}

	var facts []fact.Fact
	for _, it := range items {
		newFact := fact.Fact{
			Image:       it.PrimaryImage,
			Url:         it.Url,
			Description: it.FactText,
		}
		facts = append(facts, newFact)
	}

	return facts, nil
}
