package chuper

// Cache is a contract for cache backends implementations
type Cache interface {
	// Get returns single item from the backend if the requested item is not
	// found, returns NotFound err
	Get(key string) (interface{}, error)

	// Set sets a single item to the backend
	Set(key string, value interface{}) error

	// SetNX sets a single item to the backend, if it does not already exist
	SetNX(key string, value interface{}) (bool, error)

	// Delete deletes single item from backend
	Delete(key string) error
}
