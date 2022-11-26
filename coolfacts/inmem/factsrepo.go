package inmem

import (
	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type factsRepo struct {
	facts []coolfact.Fact
}

func NewFactsRepository() *factsRepo {
	return &factsRepo{
		// TODO: add createdAt to the facts.
		facts: []coolfact.Fact{
			{
				Image:       "https://images2.minutemediacdn.com/image/upload/v1556645500/shape/cover/entertainment/D5aliXvWsAEcYoK-fe997566220c082b98030508e654948e.jpg",
				Description: "Did you know sonic is a hedgehog?!",
			},
			{
				Image:       "https://images2.minutemediacdn.com/image/upload/v1556641470/shape/cover/entertainment/uncropped-Screen-Shot-2019-04-30-at-122411-PM-3b804f143c543dfab4b75c81833bed1b.jpg",
				Description: "You won't believe what happened to Arya!",
			},
		},
	}
}

func (r *factsRepo) GetFacts() ([]coolfact.Fact, error) {
	//TODO: before returning the facts, sort the facts according to the createdAt
	// Use sort.Sort method with the slice byCreatedAt. The most recent facts will be return first.
	// For explain on sort.Sort, see example: https://gobyexample.com/sorting-by-functions
	// Check methods of time.Time to decide how can you check if one time is after another one. (https://pkg.go.dev/time#Time.Before)

	return r.facts, nil
}

func (r *factsRepo) CreateFact(fct coolfact.Fact) error {
	// TODO: implement
}

type byCreatedAt []coolfact.Fact

// TODO: make type byCreatedAt implement sort.Interface. Example: https://gobyexample.com/sorting-by-functions
