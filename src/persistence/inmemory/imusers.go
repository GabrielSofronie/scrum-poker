package inmemory

import (
	"entities"
	"fmt"
	"persistence"
	"repositories"
)

// Users collection
type Users struct {
	persistence.Datastorer
	Collection map[string]entities.User
}

// NewUsers creates a UserUserRepository
func NewUsers(ds persistence.Datastorer) repositories.UserRepository {
	return &Users{ds,
		make(map[string]entities.User),
	}
}

// Create users
func (imu *Users) Create(usr entities.User) error {
	imu.Collection[usr.ID()] = usr
	return nil
}

// FindByID returns user with provided id
func (imu Users) FindByID(id string) (entities.User, error) {
	if _, ok := imu.Collection[id]; !ok {
		return entities.User{}, fmt.Errorf("Cannot find user with ID: %s", id)
	}
	return imu.Collection[id], nil
}
