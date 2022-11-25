package inmem

import (
	"github.com/FTBpro/go-workshop/coolfacts/fact"
)

type factsRepo struct {
	facts []fact.Fact
}

func NewFactsRepository() *factsRepo {
	// TODO: init facts repo
}

func (r *factsRepo) GetFacts() ([]fact.Fact, error) {
	// TODO: implement
}
