package store

import "context"

// Service represents a service that interacts with a database.
type Service interface {
	// * HERE we can add things like CRUD operations

	// Risks
	// Risk RiskService

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}

// CRUD operations for the Risk entity
type RiskService interface {

	Query(ctx context.Context) error

}