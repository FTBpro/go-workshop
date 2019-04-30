package inmem

import (
	"fmt"
	"github.com/FTBpro/go-workshop/coolfacts/fact"
)

const keyFormat = "image=%v,url=%v,description=%v"

type factStore struct {
	data   []fact.Fact
	hashes map[string]bool
	index  int
}

func NewFactStore() *factStore {
	return &factStore{
		data:   make([]fact.Fact, 0),
		hashes: make(map[string]bool),
		index:  0,
	}
}

func (s factStore) Get(i int) fact.Fact {
	if i >= len(s.data) {
		return fact.Fact{}
	}
	return s.data[i]
}

func (s *factStore) GetNext() fact.Fact {
	if len(s.data) == 0 {
		return fact.Fact{}
	}
	value := s.data[s.index]
	s.index++
	return value
}

func (s *factStore) Append(fact fact.Fact) int {
	key := s.generateKey(fact)
	if !s.hashes[key] {
		s.data = append(s.data, fact)
		s.hashes[key] = true
	}
	return len(s.data) - 1
}

// private

func (s *factStore) generateKey(f fact.Fact) string {
	return fmt.Sprintf(keyFormat, f.Image, f.Url, f.Description)
}
