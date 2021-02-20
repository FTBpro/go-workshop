package fact

type Fact struct {
	Image       string
	Description string
}

type Repository struct {
	facts []Fact
}

func (r *Repository) Add(f Fact) {
	// append is a Go built in function
	r.facts = append([]Fact{f}, r.facts...)
}

func (r *Repository) GetAll() []Fact {
	return r.facts
}
