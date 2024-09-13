package store

// Service represents a service that interacts with a database.
type Service interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	// TODO
	// Health() map[string]string

	// * HERE we can add things like CRUD operations

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
}
