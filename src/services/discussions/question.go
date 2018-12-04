package discussions

import (
	"entities"
	"repositories"
)

// QuestionModel request
type QuestionModel struct {
	DeveloperID string
	StoryID     string
	Content     string
}

// QuestionResponse struct
type QuestionResponse struct {
	Developer string
	Story     string
	Question  string
	Err       error
}

// Inquirer interface
type Inquirer interface {
	Execute(QuestionModel) QuestionResponse
}

type inquirer struct {
	users   repositories.UserRepository
	stories repositories.StoryRepository
}

// NewInquirer returns an Inquirer
func NewInquirer(userRepository repositories.UserRepository, storyRepository repositories.StoryRepository) Inquirer {
	return inquirer{userRepository, storyRepository}
}

func (i inquirer) Execute(qm QuestionModel) QuestionResponse {
	dev, err := i.users.FindByID(qm.DeveloperID)
	if err != nil {
		return QuestionResponse{Err: err}
	}

	story, err := i.stories.WithID(qm.StoryID)
	if err != nil {
		return QuestionResponse{Err: err}
	}

	quest := entities.NewQuestion(qm.DeveloperID, qm.StoryID, qm.Content)
	if err = i.stories.UpdateDiscussion(story.ID(), quest); err != nil {
		return QuestionResponse{Err: err}
	}

	return QuestionResponse{
		Developer: dev.Name(),
		Story:     story.Title(),
		Question:  qm.Content,
	}
}
