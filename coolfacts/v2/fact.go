package main

type fact struct {
	Image       string
	Url         string
	Description string
}

type store struct {
	facts []fact
}

func (s *store) add(f fact) {
	// append is a Go built in function
	s.facts = append(s.facts, f)
}

func (s store) getAll() []fact {
	return s.facts

}
