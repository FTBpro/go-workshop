package inmem

import (
	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type factsRepo struct {
	facts []coolfact.Fact
}

func NewFactsRepository() *factsRepo {
	return &factsRepo{
		facts: []coolfact.Fact{
			{
				Topic:       "Games",
				Description: "Did you know sonic is a hedgehog?!",
			},
			{
				Topic:       "TV",
				Description: "You won't believe what happened to Arya!",
			},
		},
	}
}

func (r *factsRepo) GetFacts() ([]coolfact.Fact, error) {
	return r.facts, nil
}
