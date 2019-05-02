package fact

type Fact struct {
	Image       string
	Url         string
	Description string
}

type Store struct {
	Facts []Fact
}

func (s *Store) Add(f Fact) {
	// append is a Go built in function
	s.Facts = append(s.Facts, f)
}

func (s Store) GetAll() []Fact {
	return s.Facts
}
