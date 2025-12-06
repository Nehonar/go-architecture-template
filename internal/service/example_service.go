package service

import (
	"context"
)

// ExampleService defines a simple business service.
// Replace this with real services in each new project.
type ExampleService interface {
	DoSomething(ctx context.Context) error
}

// ExampleService is the concrete implementation of ExampleService.
type exampleService struct{}

// NewExampleService returns a new ExampleService.
func NewExampleService() ExampleService {
	return &exampleService{}
}

// DoSomething executes a simple business operation.
// Real projects will replace this with actual domain logic.
func (s *exampleService) DoSomething(ctx context.Context) error {
	// TODO: implement business logic here
	return nil
}
