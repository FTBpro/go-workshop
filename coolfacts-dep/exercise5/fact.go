package main

type fact struct {
	Image       string
	Description string
}

type repo struct {
	facts []fact
}

func (s *repo) add(f fact) {
	// append is a Go built in function
	s.facts = append(s.facts, f)
}

func (s *repo) getAll() []fact {
	return s.facts
}
