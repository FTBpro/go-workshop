package fact

type Fact struct {
	Image       string
	Url         string
	Description string
}

type Store struct {
	facts []Fact
}

func (s *Store) Add(f Fact) {
	// append is a Go built in function
	s.facts = append(s.facts, f)
}

func (s Store) GetAll() []Fact {
	return s.facts
}
