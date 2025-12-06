package postgres

// ExampleRepository defines how a concrete storage layer should behave.
// Replace this with real repository interfaces in actual projects.
type ExampleRepository interface {
	// TODO: Define repository methods like:
	// Save(example *domain.Example) error
	// FindByID(id int) (*domain.Example, error)
}

func NewExampleRepository() ExampleRepository {
	// TODO: return a struct implementing the interface
	return nil
}
