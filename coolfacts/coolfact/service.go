package coolfact

type Repository interface {
	// TODO: add functions decleration
	// - getFacts. Returns a slice of Fact and an error
}

type service struct {
	// TODO: add field factsRepo
}

func NewService(factsRepo Repository) *service {
	// TODO: init a new service with factsRepo
}

func (s *service) GetFacts() ([]Fact, error) {
	// TODO: implement getFacts, using the factsRepo
}
