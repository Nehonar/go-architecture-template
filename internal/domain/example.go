package domain

// Example represents a domain entity.

// Domain type model the core buisness concepts to the system.

// They are independent from infraestructure concerns such as HTTP,
// databases, serialization or external services.

// In real projects, replace this with domain models like:
// - User
// - Product
// - Order
// - ShardNode

// Domain models express the vocabulary of the business and are the
// most stable and long-lived part of the architecture.

type Example struct {
	ID   int
	Name string
}
