package inmem

import (
	"sort"
	
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
	sort.Sort(byCreatedAt(r.facts))
	
	return r.facts, nil
}

func (r *factsRepo) CreateFact(fact coolfact.Fact) error {
	r.facts = append(r.facts, fact)
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
