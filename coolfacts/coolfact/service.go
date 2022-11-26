package coolfact

import (
	"fmt"
	"time"
)

type Repository interface {
	GetFacts() ([]Fact, error)
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

func (s *service) GetFacts() ([]Fact, error) {
	facts, err := s.factsRepo.GetFacts()
	if err != nil {
		return nil, fmt.Errorf("factsService.GetFacts: %w", err)
	}

	return facts, nil
}

func (s *service) CreateFact(fct Fact) error {
	fct.CreatedAt = time.Now()
	if err := s.factsRepo.CreateFact(fct); err != nil {
		return fmt.Errorf("factsService.CreateFact: %w", err)
	}

	return nil
}
