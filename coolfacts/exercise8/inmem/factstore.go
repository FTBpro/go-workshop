package inmem

import (
	"github.com/FTBpro/go-workshop/coolfacts/exercise8/fact"
)


type FactStore struct {
	facts []fact.Fact
}

func (s *FactStore) Add(f fact.Fact) {
	// append is a Go built in function
	s.facts = append(s.facts, f)
}

func (s *FactStore) GetAll() []fact.Fact {
	return s.facts
}

