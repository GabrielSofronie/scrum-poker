package tests

import (
	"entities"
	"services/discussions"
	"testing"
)

func TestDeveloperAsksQuestion(t *testing.T) {
	setup()

	// Add test developer and story
	testDeveloper := entities.NewDeveloper("test-dev")
	userRepository.Create(testDeveloper)

	testStory := entities.NewStory("test-story", "This is a story")
	storyRepository.Create(testStory)

	// Ask
	inquirer := discussions.NewInquirer(userRepository, storyRepository)
	quest := "Is this a question?"
	questModel := discussions.QuestionModel{
		DeveloperID: testDeveloper.ID(),
		StoryID:     testStory.ID(),
		Content:     quest,
	}
	response := inquirer.Execute(questModel)

	if response.Err != nil {
		t.Errorf(response.Err.Error())
	}

	if response.Developer != testDeveloper.Name() {
		t.Errorf("Expected to have developer name: [%s] , got [%s] ", testDeveloper.Name(), response.Developer)
	}

	if response.Story != testStory.Title() {
		t.Errorf("Expected to have story title: [%s] , got [%s] ", testStory.Title(), response.Story)
	}

	if response.Question != quest {
		t.Errorf("Expected to have question: [%s] , got [%s] ", quest, response.Question)
	}
}

func TestOwnerAnswersQuestion(t *testing.T) {
	setup()

	// Add test owner and story
	testOwner := entities.NewOwner("test-owner")
	userRepository.Create(testOwner)

	testStory := entities.NewStory("test-story", "This is a story")
	testQuestion := entities.NewQuestion(testOwner.ID(), testStory.ID(), "How?")
	testStory.AddDiscussion(testQuestion)
	storyRepository.Create(testStory)

	// Answer
	answer := discussions.NewAnswer(userRepository, storyRepository)
	answerModel := discussions.AnswerModel{
		OwnerID:   testOwner.ID(),
		StoryID:   testStory.ID(),
		ReplyToID: testQuestion.ID(),
		Content:   "This is an answer",
	}
	response := answer.Execute(answerModel)

	if response.Err != nil {
		t.Errorf(response.Err.Error())
	}

	if response.Owner != testOwner.Name() {
		t.Errorf("Expected to have owner name: [%s] , got [%s] ", testOwner.Name(), response.Owner)
	}

	if response.Question != testQuestion.Content() {
		t.Errorf("Expected to have questions: [%s], got [%s]", testQuestion.Content(), response.Question)
	}

	if response.Answer != answerModel.Content {
		t.Errorf("Expected to have answer: [%s], got [%s]", answerModel.Content, response.Answer)
	}
}
