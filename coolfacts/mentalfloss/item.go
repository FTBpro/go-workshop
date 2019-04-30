package mentalfloss

import "github.com/FTBpro/go-workshop/coolfacts/fact"

type item struct {
	Url          string `json:"url"`
	FactText     string `json:"fact"`
	PrimaryImage string `json:"primaryImage"`
}

func (it item) ToFact() fact.Fact {
	return fact.Fact{
		Image:       it.PrimaryImage,
		Url:         it.Url,
		Description: it.FactText,
	}
}
