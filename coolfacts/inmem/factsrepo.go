package inmem

import (
	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type factsRepo struct {
	facts []coolfact.Fact
}

func NewFactsRepository() *factsRepo {
	// TODO: init facts repo
}

func (r *factsRepo) GetFacts() ([]coolfact.Fact, error) {
	// TODO: implement
}
