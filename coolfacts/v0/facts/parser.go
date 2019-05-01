package facts

import "encoding/json"

type MentalFlossItem struct {
	Id            string   `json:"id"`
	Url           string   `json:"url"`
	FactId        string   `json:"factId"`
	Headline      string   `json:"headline"`
	ShortHeadline string   `json:"shortHeadline"`
	FactText      string   `json:"fact"`
	FullStoryUrl  string   `json:"fullStoryUrl"`
	Tags          []string `json:"tags"`
	PrimaryImage  string   `json:"primaryImage"`
	ImageCredit   string   `json:"imageCredit"`
}

type Fact struct {
	Image       string
	Url         string
	Description string
}

type parser struct {
}

func NewParser() *parser {
	return &parser{}
}

func (p *parser) ParseFromPolling(b []byte) ([]Fact, error) {
	data := make([]MentalFlossItem, 0)
	if err := json.Unmarshal(b, &data); err != nil {
		return nil, err
	}
	facts := make([]Fact, 0)
	for _, fact := range data {
		parsedFact := p.parse(fact)
		facts = append(facts, parsedFact)
	}
	return facts, nil
}

func (p *parser) parse(mfi MentalFlossItem) Fact {
	return Fact{
		Image:       mfi.PrimaryImage,
		Url:         mfi.Url,
		Description: mfi.FactText,
	}
}
