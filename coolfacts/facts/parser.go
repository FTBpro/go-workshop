package facts

import "encoding/json"

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

func (p *parser) ParseFromCreate(b []byte) (Fact, error) {
	data := MentalFlossItem{}
	if err := json.Unmarshal(b, &data); err != nil {
		return Fact{}, nil
	}
	return p.parse(data), nil
}


func (p *parser) parse(mfi MentalFlossItem) Fact {
	return Fact{
		Image:       mfi.PrimaryImage,
		Url:         mfi.Url,
		Description: mfi.FactText,
	}
}
