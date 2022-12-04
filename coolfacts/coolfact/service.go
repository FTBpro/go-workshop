package coolfact

import "fmt"

type Repository interface {
	SearchFacts(filters Filters) ([]Fact, error)
	CreateFact(fct Fact) error
}

type service struct {
	factsRepo Repository
}

func NewService(factsRepo Repository) *service {
	return &service{
		factsRepo: factsRepo,
	}
}

func (s *service) SearchFacts(filters Filters) ([]Fact, error) {
	facts, err := s.factsRepo.SearchFacts(filters)
	if err != nil {
		return nil, fmt.Errorf("factsService.SearchFacts: %w", err)
	}

	return facts, nil
}

func (s *service) CreateFact(fact Fact) error {
	if err := s.factsRepo.CreateFact(fact); err != nil {
		return fmt.Errorf("factsService.CreateFact: %w", err)
	}

	return nil
}
