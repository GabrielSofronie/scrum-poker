package repositories

import (
	"entities"
)

// StoryRepository interface
type StoryRepository interface {
	Create(*entities.Story) error
	UpdateEstimation(string, entities.Estimation) error
	UpdateDiscussion(string, entities.Discussion) error
	WithID(string) (*entities.Story, error)
	List() ([]*entities.Story, error)
}
