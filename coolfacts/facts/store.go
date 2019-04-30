package facts

import "fmt"

const keyFormat = "image=%v,url=%v,description=%v"

type store struct {
	data   []Fact
	hashes map[string]bool
	index  int
}

func NewStore() *store {
	return &store{
		data:   make([]Fact, 0),
		hashes: make(map[string]bool),
		index:  0,
	}
}

func (s store) Get(i int) Fact {
	if i >= len(s.data) {
		return Fact{}
	}
	return s.data[i]
}

func (s *store) GetNext() Fact {
	if len(s.data) == 0 {
		return Fact{}
	}
	value := s.data[s.index]
	s.index++
	return value
}

func (s *store) AppendFact(fact Fact) int {
	key := s.generateKey(fact)
	if !s.hashes[key] {
		s.data = append(s.data, fact)
		s.hashes[key] = true
	}
	return len(s.data) - 1
}

func (s *store) generateKey(fact Fact) string {
	return fmt.Sprintf(keyFormat, fact.Image, fact.Url, fact.Description)
}
