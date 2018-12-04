package inmemory

import "persistence"

// InMemory store
type InMemory struct{}

// Open an InMemory connection
func (m *InMemory) Open(connect ...string) (persistence.Datastorer, error) {
	return m, nil
}
