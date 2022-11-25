package fact

import (
	"testing"
)

func TestRepo_Add(t *testing.T) {
	repo := Repository{}
	fact := Fact{
		Image:       "img",
		Description: "desc",
	}
	repo.Add(fact)

	gotFacts := repo.GetAll()
	var hasItem bool
	for _, gotFact := range gotFacts {
		if gotFact.Image == fact.Image && gotFact.Description == fact.Description {
			hasItem = true
		}
	}
	if !hasItem {
		t.Errorf("can't find added fact")
	}
}
