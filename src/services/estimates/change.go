package estimates

import (
	"entities"
	"repositories"
)

// ChangeModel request
type ChangeModel struct {
	DeveloperID string
	StoryID     string
	Value       int
}

// ChangeResponse struct
type ChangeResponse struct {
	Developer string
	Story     string
	Estimate  int
	Err       error
}

// Changer interface
type Changer interface {
	Execute(ChangeModel) ChangeResponse
}

type change struct {
	users   repositories.UserRepository
	stories repositories.StoryRepository
}

// NewChanger returns a Changer
func NewChanger(userRepository repositories.UserRepository, storyRepository repositories.StoryRepository) Changer {
	return change{userRepository, storyRepository}
}

func (c change) Execute(request ChangeModel) ChangeResponse {
	usr, err := c.users.FindByID(request.DeveloperID)
	if err != nil {
		return ChangeResponse{Err: err}
	}

	s, err := c.stories.WithID(request.StoryID)
	if err != nil {
		return ChangeResponse{Err: err}
	}

	estimation := entities.NewEstimation(usr.ID(), s.ID(), request.Value)
	if err := c.stories.UpdateEstimation(s.ID(), estimation); err != nil {
		return ChangeResponse{Err: err}
	}
	return ChangeResponse{
		Developer: usr.Name(),
		Story:     s.Title(),
		Estimate:  request.Value,
	}
}
