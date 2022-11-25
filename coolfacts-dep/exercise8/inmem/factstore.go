package inmem

import (
	"github.com/FTBpro/go-workshop/coolfacts/exercise8/fact"
)

type factRepository struct {
	facts []fact.Fact
}

func NewFactRepository() *factRepository {
	return &factRepository{}
}

func (r *factRepository) Add(f fact.Fact) {
	// append is a Go built in function
	r.facts = append(r.facts, f)
}

func (r *factRepository) GetAll() []fact.Fact {
	return r.facts
}
