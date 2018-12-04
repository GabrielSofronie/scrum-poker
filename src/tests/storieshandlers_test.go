package tests

import (
	"api/handlers"
	"entities"
	"fmt"
	"net/http"
	"net/http/httptest"
	"services/stories"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestStoryCreateHandler(t *testing.T) {
	setup()

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	storyCreator := stories.NewCreator(storyRepository)
	storyHdlr := handlers.NewStoryHandler(router, storyCreator, nil)
	storyHdlr.RegisterHandlers()

	title := "story-title"
	payload := fmt.Sprintf(`{
			"Title": "%s",
			"Description": "%s"
		}`, title, "This is a story")

	req, err := http.NewRequest("POST", "/v1/stories", strings.NewReader(payload))
	if err != nil {
		t.Error(err)
	}

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected to have HTTP Code [%d] , got [%d] ", http.StatusCreated, recorder.Code)
	}
}

func TestStoryListHandler(t *testing.T) {
	setup()

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	storyTitle := "test-story"
	testStory := entities.NewStory(storyTitle, "This is a story")
	storyRepository.Create(testStory)

	storyLister := stories.NewLister(storyRepository)
	storyHdlr := handlers.NewStoryHandler(router, nil, storyLister)
	storyHdlr.RegisterHandlers()

	req, err := http.NewRequest("GET", "/v1/stories", nil)
	if err != nil {
		t.Error(err)
	}

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected to have HTTP Code [%d] , got [%d] ", http.StatusOK, recorder.Code)
	}
}
