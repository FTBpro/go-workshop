package inmem

import (
	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type factsRepo struct {
	facts []coolfact.Fact
}

func NewFactsRepository(facts ...coolfact.Fact) *factsRepo {
	return &factsRepo{
		facts: facts,
	}
}

func (r *factsRepo) GetFacts() ([]coolfact.Fact, error) {
	//TODO: before returning the facts, sort the facts according to the createdAt
	// Use sort.Sort method with the slice byCreatedAt. The most recent facts will be return first.
	// For explain on sort.Sort, see example: https://gobyexample.com/sorting-by-functions
	// Check methods of time.Time to decide how can you check if one time is after another one. (https://pkg.go.dev/time#Time.Before)

	return r.facts, nil
}

func (r *factsRepo) CreateFact(fact coolfact.Fact) error {
	// TODO: implement
}

type byCreatedAt []coolfact.Fact

// TODO: make type byCreatedAt implement sort.Interface. Example: https://gobyexample.com/sorting-by-functions
