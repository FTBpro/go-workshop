package coolfact

import "fmt"

type Repository interface {
	GetFacts() ([]Fact, error)
}

type service struct {
	factsRepo Repository
}

func NewService(factsRepo Repository) *service {
	return &service{
		factsRepo: factsRepo,
	}
}

func (s *service) GetFacts() ([]Fact, error) {
	facts, err := s.factsRepo.GetFacts()
	if err != nil {
		return nil, fmt.Errorf("factsService.GetFacts: %w", err)
	}

	return facts, nil
}
