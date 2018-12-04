package stories

import (
	"entities"
	"repositories"
)

// StoryModel reuqest
type StoryModel struct {
	Title       string
	Description string
}

// Creator interface
type Creator interface {
	Execute(StoryModel) CreateResponse
}

// CreateResponse struct
type CreateResponse struct {
	Title string
	Err   error
}

type owner struct {
	repositories.StoryRepository
}

// NewCreator returns a Creator
func NewCreator(storyRepository repositories.StoryRepository) Creator {
	return owner{storyRepository}
}

func (o owner) Execute(sm StoryModel) CreateResponse {
	story := entities.NewStory(sm.Title, sm.Description)

	if err := o.Create(story); err != nil {
		return CreateResponse{Err: err}
	}

	return CreateResponse{Title: sm.Title}
}
