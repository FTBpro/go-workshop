package facts

type store struct {
	Data []Fact
}

func NewStore() *store {
	return &store{}
}

func (s *store) Get() []Fact {
	return s.Data
}

func (s *store) Set(data []Fact) {
	s.Data = data
}
