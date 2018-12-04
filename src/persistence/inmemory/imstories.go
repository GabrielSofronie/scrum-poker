package inmemory

import (
	"entities"
	"fmt"
	"persistence"
	"repositories"
)

// Stories collection
type Stories struct {
	persistence.Datastorer
	Collection map[string]*entities.Story
}

// NewStories creates a StoryRepository
func NewStories(ds persistence.Datastorer) repositories.StoryRepository {
	return &Stories{ds,
		make(map[string]*entities.Story),
	}
}

// Create stories
func (ims *Stories) Create(s *entities.Story) error {
	ims.Collection[s.ID()] = s
	return nil
}

// UpdateEstimation of story
func (ims *Stories) UpdateEstimation(id string, estimate entities.Estimation) error {
	if _, ok := ims.Collection[id]; !ok {
		return fmt.Errorf("Cannot update missing story with ID: %s", id)
	}
	ims.Collection[id].DeleteEstimate(estimate)
	ims.Collection[id].SetEstimate(estimate)
	return nil
}

// UpdateDiscussion of story
func (ims *Stories) UpdateDiscussion(id string, discuss entities.Discussion) error {
	if _, ok := ims.Collection[id]; !ok {
		return fmt.Errorf("Cannot update missing story with ID: %s", id)
	}
	ims.Collection[id].DeleteDiscussion(discuss)
	ims.Collection[id].AddDiscussion(discuss)
	return nil
}

// WithID returns user with provided id
func (ims Stories) WithID(id string) (*entities.Story, error) {
	if _, ok := ims.Collection[id]; !ok {
		return nil, fmt.Errorf("Cannot find story with ID: %s", id)
	}
	return ims.Collection[id], nil
}

// List all stories
func (ims Stories) List() ([]*entities.Story, error) {
	stories := []*entities.Story{}
	for _, s := range ims.Collection {
		stories = append(stories, s)
	}
	return stories, nil
}
