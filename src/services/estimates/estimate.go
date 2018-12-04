package estimates

import (
	"entities"
	"errors"
	"repositories"
)

// EstimateModel request
type EstimateModel struct {
	DeveloperID string
	StoryID     string
	Value       int
}

// EstimateResponse struct
type EstimateResponse struct {
	Developer string
	Story     string
	Estimate  int
	Err       error
}

// Estimator interface
type Estimator interface {
	Execute(EstimateModel) EstimateResponse
}

type estimate struct {
	users   repositories.UserRepository
	stories repositories.StoryRepository
}

// NewEstimator returns an Estimator
func NewEstimator(userRepository repositories.UserRepository, storyRepository repositories.StoryRepository) Estimator {
	return estimate{userRepository, storyRepository}
}

func (e estimate) Execute(request EstimateModel) EstimateResponse {
	usr, err := e.users.FindByID(request.DeveloperID)
	if err != nil {
		return EstimateResponse{Err: err}
	}

	if usr.Role() != entities.Developer {
		return EstimateResponse{Err: errors.New("Only developers can provide estimations")}
	}

	s, err := e.stories.WithID(request.StoryID)
	if err != nil {
		return EstimateResponse{Err: err}
	}

	estimation := entities.NewEstimation(usr.ID(), s.ID(), request.Value)
	if err := e.stories.UpdateEstimation(s.ID(), estimation); err != nil {
		return EstimateResponse{Err: err}
	}

	return EstimateResponse{
		Developer: usr.Name(),
		Story:     s.Title(),
		Estimate:  request.Value,
	}
}
