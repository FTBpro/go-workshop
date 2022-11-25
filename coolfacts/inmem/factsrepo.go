package inmem

import (
	"github.com/FTBpro/go-workshop/coolfacts/fact"
)

type factsRepo struct {
	facts []fact.Fact
}

func NewFactsRepository() *factsRepo {
	return &factsRepo{
		facts: []fact.Fact{
			{
				Image:       "https://images2.minutemediacdn.com/image/upload/v1556645500/shape/cover/entertainment/D5aliXvWsAEcYoK-fe997566220c082b98030508e654948e.jpg",
				Description: "Did you know sonic is a hedgehog?!",
			},
			{
				Image:       "https://images2.minutemediacdn.com/image/upload/v1556641470/shape/cover/entertainment/uncropped-Screen-Shot-2019-04-30-at-122411-PM-3b804f143c543dfab4b75c81833bed1b.jpg",
				Description: "You won't believe what happened to Arya!",
			},
		},
	}
}

func (r *factsRepo) GetFacts() ([]fact.Fact, error) {
	return r.facts, nil
}
