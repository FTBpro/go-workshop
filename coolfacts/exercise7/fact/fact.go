package fact

type Fact struct {
	Image       string
	Description string
}

type Store struct {
	facts []Fact
}

func (s *Store) Add(f Fact) {
	// append is a Go built in function
	s.facts = append([]Fact{f}, s.facts...)
}

func (s *Store) GetAll() []Fact {
	return s.facts
}
