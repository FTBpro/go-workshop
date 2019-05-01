package fact

import (
	"fmt"
)

type Provider interface {
	Facts() ([]Fact, error)
}

type Store interface {
	Get(i int) Fact
	GetNext() Fact
	Append(fact Fact) int
}

type service struct {
	provider Provider
	store    Store
}

func NewService(s Store, r Provider) *service {
	return &service{
		store:    s,
		provider: r,
	}
}

func (s *service) UpdateFacts() error {
	facts, err := s.provider.Facts()
	if err != nil {
		return fmt.Errorf("fact.service.UpdateFacts failed returiev facts %v", err)
	}

	for _, fact := range facts {
		s.store.Append(fact)
	}

	return nil
}
