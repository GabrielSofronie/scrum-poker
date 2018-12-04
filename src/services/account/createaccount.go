package account

import (
	"entities"
	"repositories"
)

// CreateModel request
type CreateModel struct {
	Name string
	Role entities.Role
}

// Maker interface
type Maker interface {
	Execute(CreateModel) CreateResponse
}

// CreateResponse struct
type CreateResponse struct {
	ID   string
	Name string
	Err  error
}

type maker struct {
	repositories.UserRepository
}

func (m maker) Execute(account CreateModel) CreateResponse {
	usr := createUser(account.Name, account.Role)
	if err := m.Create(usr); err != nil {
		return CreateResponse{Err: err}
	}
	return CreateResponse{
		ID:   usr.ID(),
		Name: usr.Name(),
	}
}

// NewMaker returns AccountMaker interface
func NewMaker(repo repositories.UserRepository) Maker {
	return maker{repo}
}

func createUser(name string, role entities.Role) entities.User {
	if role == entities.Developer {
		return entities.NewDeveloper(name)
	}

	if role == entities.Moderator {
		return entities.NewModerator(name)
	}

	return entities.NewOwner(name)
}
