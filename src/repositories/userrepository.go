package repositories

import "entities"

// UserRepository interface
type UserRepository interface {
	Create(entities.User) error
	FindByID(string) (entities.User, error)
}
