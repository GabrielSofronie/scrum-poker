package tests

import (
	"api/handlers"
	"entities"
	"fmt"
	"net/http"
	"net/http/httptest"
	"services/discussions"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestStoryQuestionHandler(t *testing.T) {
	setup()

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Add test developer and story
	testDeveloper := entities.NewDeveloper("test-dev")
	userRepository.Create(testDeveloper)

	testStory := entities.NewStory("test-story", "This is a story")
	storyRepository.Create(testStory)

	inquirer := discussions.NewInquirer(userRepository, storyRepository)
	discussionHdlr := handlers.NewDiscussionHandler(router, inquirer, nil)
	discussionHdlr.RegisterHandlers()

	content := "What is the purpose?"
	payload := fmt.Sprintf(`{
			"DeveloperId": "%s",
			"Content": "%s"
		}`, testDeveloper.ID(), content)

	path := fmt.Sprintf("/v1/stories/%s/comment", testStory.ID())
	req, err := http.NewRequest("POST", path, strings.NewReader(payload))
	if err != nil {
		t.Error(err)
	}

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected to have HTTP Code [%d] , got [%d] ", http.StatusCreated, recorder.Code)
	}
}

func TestStoryReplyHandler(t *testing.T) {
	setup()

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Add test owner and story
	testOwner := entities.NewOwner("test-owner")
	userRepository.Create(testOwner)

	testStory := entities.NewStory("test-story", "This is a story")
	testQuestion := entities.NewQuestion(testOwner.ID(), testStory.ID(), "How to do it?")
	testStory.AddDiscussion(testQuestion)
	storyRepository.Create(testStory)

	answer := discussions.NewAnswer(userRepository, storyRepository)
	discussionHdlr := handlers.NewDiscussionHandler(router, nil, answer)
	discussionHdlr.RegisterHandlers()

	content := "As easy as possible"
	payload := fmt.Sprintf(`{
			"OwnerId": "%s",
			"ReplyToID": "%s",
			"Content": "%s"
		}`, testOwner.ID(), testQuestion.ID(), content)

	path := fmt.Sprintf("/v1/stories/%s/comment", testStory.ID())
	req, err := http.NewRequest("PUT", path, strings.NewReader(payload))
	if err != nil {
		t.Error(err)
	}

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected to have HTTP Code [%d] , got [%d] ", http.StatusCreated, recorder.Code)
	}
}
