package inmem

import (
	"sort"

	"github.com/FTBpro/go-workshop/coolfacts/coolfact"
)

type factsRepo struct {
	factsByTopic map[string][]coolfact.Fact
}

func NewFactsRepository(facts ...coolfact.Fact) *factsRepo {
	factsByTopic := map[string][]coolfact.Fact{}
	for _, fact := range facts {
		factsByTopic[fact.Topic] = append(factsByTopic[fact.Topic], fact)
	}

	return &factsRepo{
		factsByTopic: factsByTopic,
	}
}

func (r *factsRepo) GetFacts(filters coolfact.Filters) ([]coolfact.Fact, error) {
	var facts []coolfact.Fact
	if filters.Topic != "" {
		facts = r.factsByTopic[filters.Topic]
	} else {
		facts = r.allFacts()
	}

	sort.Sort(byCreatedAt(facts))

	if filters.Limit < len(facts) {
		facts = facts[:filters.Limit]
	}

	return facts, nil
}

func (r *factsRepo) CreateFact(fact coolfact.Fact) error {
	r.factsByTopic[fact.Topic] = append(r.factsByTopic[fact.Topic], fact)

	return nil
}

func (s *factsRepo) allFacts() []coolfact.Fact {
	var allFacts []coolfact.Fact
	for _, facts := range s.factsByTopic {
		allFacts = append(allFacts, facts...)
	}

	return allFacts
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
