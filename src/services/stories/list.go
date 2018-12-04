package stories

import (
	"entities"
	"repositories"
)

// Lister interface
type Lister interface {
	Execute() ListResponse
}

// ListResponse struct
type ListResponse struct {
	Titles []string
	Err    error
}

type lister struct {
	repositories.StoryRepository
}

// NewLister returns a StoryLister
func NewLister(storyRepository repositories.StoryRepository) Lister {
	return lister{storyRepository}
}

func (l lister) Execute() ListResponse {
	stories, err := l.List()
	if err != nil {
		return ListResponse{Err: err}
	}

	return ListResponse{
		Titles: titlesOf(stories),
	}
}

func titlesOf(stories []*entities.Story) []string {
	var titles []string
	for _, s := range stories {
		titles = append(titles, s.Title())
	}
	return titles
}
