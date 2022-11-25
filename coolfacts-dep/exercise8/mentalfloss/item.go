package mentalfloss

import "github.com/FTBpro/go-workshop/coolfacts/exercise8/fact"

type item struct {
	FactText     string `json:"fact"`
	PrimaryImage string `json:"primaryImage"`
}

func (it item) ToFact() fact.Fact {
	return fact.Fact{
		Image:       it.PrimaryImage,
		Description: it.FactText,
	}
}
