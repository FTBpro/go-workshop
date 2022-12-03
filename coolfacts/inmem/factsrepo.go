package inmem

import (
	"sort"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type factsRepo struct {
	factsByTopic map[string][]coolfact.Fact
}

func NewFactsRepository(facts ...coolfact.Fact) *factsRepo {
	// TODO: fix initialization according to the new field type
	return &factsRepo{
		facts: facts,
	}
}

func (r *factsRepo) GetFacts(filters coolfact.Filters) ([]coolfact.Fact, error) {
	// TODO: fix method. Return according to the filters.
	// note - topic is optional.
	sort.Sort(byCreatedAt(r.facts))

	return r.facts, nil
}

func (r *factsRepo) CreateFact(fact coolfact.Fact) error {
	// TODO: fix according to the new field type
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
