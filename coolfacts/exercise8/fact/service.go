package fact

import (
	"fmt"
)

type Provider interface {
	Facts() ([]Fact, error)
}

type Repository interface {
	Add(f Fact)
	GetAll() []Fact
}

type service struct {
	provider Provider
	repo     Repository
}

func NewService(p Provider, s Repository) *service {
	return &service{
		provider: p,
		repo:     s,
	}
}

func (s *service) UpdateFacts() error {
	facts, err := s.provider.Facts()
	if err != nil {
		return fmt.Errorf("fact.service.UpdateFacts failed returiev facts %v", err)
	}

	for _, fact := range facts {
		s.repo.Add(fact)
	}

	return nil
}
