package inmem

import (
	"sort"
	"time"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type factsRepo struct {
	facts []coolfact.Fact
}

func NewFactsRepository() *factsRepo {
	return &factsRepo{
		facts: []coolfact.Fact{
			{
				Image:       "https://images2.minutemediacdn.com/image/upload/v1556645500/shape/cover/entertainment/D5aliXvWsAEcYoK-fe997566220c082b98030508e654948e.jpg",
				Description: "Did you know sonic is a hedgehog?!",
				CreatedAt:   time.Now(),
			},
			{
				Image:       "https://images2.minutemediacdn.com/image/upload/v1556641470/shape/cover/entertainment/uncropped-Screen-Shot-2019-04-30-at-122411-PM-3b804f143c543dfab4b75c81833bed1b.jpg",
				Description: "You won't believe what happened to Arya!",
				CreatedAt:   time.Now(),
			},
		},
	}
}

func (r *factsRepo) GetFacts() ([]coolfact.Fact, error) {
	sort.Sort(byCreatedAt(r.facts))

	return r.facts, nil
}

func (r *factsRepo) CreateFact(fct coolfact.Fact) error {
	r.facts = append(r.facts, fct)
	return nil
}

type byCreatedAt []coolfact.Fact

func (s byCreatedAt) Len() int {
	return len(s)
}
func (s byCreatedAt) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byCreatedAt) Less(i, j int) bool {
	return s[i].CreatedAt.After(s[j].CreatedAt)
}
