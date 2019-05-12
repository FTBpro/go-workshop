package fact

import (
	"testing"
)

func TestStore_Add(t *testing.T) {
	store := Store{}
	fact := Fact{
		Image:       "img",
		Url:         "url",
		Description: "desc",
	}
	store.Add(fact)

	gotFacts := store.GetAll()
	var hasItem bool
	for _, gotFact := range gotFacts {
		if gotFact.Url == fact.Url && gotFact.Image == fact.Image && gotFact.Description == fact.Description {
			hasItem = true
		}
	}
	if !hasItem {
		t.Errorf("can't find added fact")
	}
}
