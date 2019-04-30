package fact

import (
	"fmt"
)

type Retriever interface {
	Facts() ([]Fact, error)
}

type Store interface {
	Get(i int) Fact
	GetNext() Fact
	Append(fact Fact) int
}

type service struct {
	store     Store
	retriever Retriever
}

func NewService(s Store, r Retriever) *service {
	return &service{
		store:     s,
		retriever: r,
	}
}

func (s *service) UpdateFacts() error {
	facts, err := s.retriever.Facts()
	if err != nil {
		return fmt.Errorf("fact.service.UpdateFacts failed returiev facts %v", err)
	}

	for _, fact := range facts {
		s.store.Append(fact)
	}

	return nil
}
