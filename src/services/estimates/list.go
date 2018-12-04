package estimates

import (
	"entities"
	"repositories"
)

// ListModel request
type ListModel struct {
	StoryID string
}

// ListResponse struct
type ListResponse struct {
	Estimates map[string]int
	Err       error
}

// Lister interface
type Lister interface {
	Execute(ListModel) ListResponse
}

type list struct {
	stories repositories.StoryRepository
}

// NewLister returns a Lister
func NewLister(storyRepository repositories.StoryRepository) Lister {
	return list{storyRepository}
}

func (l list) Execute(request ListModel) ListResponse {
	s, err := l.stories.WithID(request.StoryID)
	if err != nil {
		return ListResponse{Err: err}
	}

	return ListResponse{
		Estimates: estimatesOf(s),
	}
}

func estimatesOf(s *entities.Story) map[string]int {
	estimates := make(map[string]int)
	for _, e := range s.Estimates() {
		estimates[e.DeveloperID()] = e.Value()
	}
	return estimates
}
