package discussions

import (
	"entities"
	"repositories"
)

// AnswerModel request
type AnswerModel struct {
	OwnerID   string
	StoryID   string
	ReplyToID string
	Content   string
}

type AnswerResponse struct {
	Owner    string
	Question string
	Answer   string
	Err      error
}

// Answer interface
type Answer interface {
	Execute(AnswerModel) AnswerResponse
}

type answer struct {
	users   repositories.UserRepository
	stories repositories.StoryRepository
}

// NewAnswer returns an Answer interface
func NewAnswer(userRepository repositories.UserRepository, storyRepository repositories.StoryRepository) Answer {
	return answer{
		users:   userRepository,
		stories: storyRepository,
	}
}

func (a answer) Execute(answerModel AnswerModel) AnswerResponse {
	owner, err := a.users.FindByID(answerModel.OwnerID)
	if err != nil {
		return AnswerResponse{Err: err}
	}

	story, err := a.stories.WithID(answerModel.StoryID)
	if err != nil {
		return AnswerResponse{Err: err}
	}

	an := entities.NewAnswer(answerModel.OwnerID, answerModel.StoryID, answerModel.ReplyToID, answerModel.Content)
	if err = a.stories.UpdateDiscussion(answerModel.StoryID, an); err != nil {
		return AnswerResponse{Err: err}
	}

	return AnswerResponse{
		Owner:    owner.Name(),
		Question: questionContent(answerModel.ReplyToID, story.Discussions()),
		Answer:   answerModel.Content,
	}
}

func questionContent(id string, discussions []*entities.Discussion) string {
	for _, d := range discussions {
		if d.ID() == id {
			return d.Content()
		}
	}

	return ""
}
