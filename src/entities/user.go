package entities

import (
	"workmate"
)

// User entity
type User struct {
	id   string
	name string
	role Role
}

// ID of the user
func (u User) ID() string {
	return u.id
}

// Name of the user
func (u User) Name() string {
	return u.name
}

// Role of the user
func (u User) Role() Role {
	return u.role
}

// NewDeveloper returns a developer user
func NewDeveloper(name string) User {
	return User{
		id:   workmate.UUID(),
		name: name,
		role: Developer,
	}
}

// NewModerator returns a moderator user
func NewModerator(name string) User {
	return User{
		id:   workmate.UUID(),
		name: name,
		role: Moderator,
	}
}

// NewOwner returns a product owner user
func NewOwner(name string) User {
	return User{
		id:   workmate.UUID(),
		name: name,
		role: Owner,
	}
}
