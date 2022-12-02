package inmem

import (
	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type factsRepo struct {
	facts []coolfact.Fact
}

func NewFactsRepository(facts ...coolfact.Fact) *factsRepo {
	return &factsRepo{
		facts: facts,
	}
}

func (r *factsRepo) GetFacts() ([]coolfact.Fact, error) {
	return r.facts, nil
}
