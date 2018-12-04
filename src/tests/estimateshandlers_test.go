package tests

import (
	"api/handlers"
	"entities"
	"fmt"
	"net/http"
	"net/http/httptest"
	"services/estimates"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

func TestStoryEstimateHandler(t *testing.T) {
	setup()

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Add test developer and story
	testDeveloper := entities.NewDeveloper("test-dev")
	userRepository.Create(testDeveloper)

	testStory := entities.NewStory("test-story", "This is a story")
	storyRepository.Create(testStory)

	estimator := estimates.NewEstimator(userRepository, storyRepository)
	estimateHdlr := handlers.NewEstimateHandler(router, estimator, nil, nil)
	estimateHdlr.RegisterHandlers()

	estimate := 5
	payload := fmt.Sprintf(`{
			"DeveloperId": "%s",
			"Value": %d
		}`, testDeveloper.ID(), estimate)

	path := fmt.Sprintf("/v1/stories/%s/estimates", testStory.ID())
	req, err := http.NewRequest("POST", path, strings.NewReader(payload))
	if err != nil {
		t.Error(err)
	}

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Errorf("Expected to have HTTP Code [%d] , got [%d] ", http.StatusCreated, recorder.Code)
	}
}

func TestChangeEstimateHandler(t *testing.T) {
	setup()

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	// Add test developer and story
	testDeveloper := entities.NewDeveloper("test-dev")
	userRepository.Create(testDeveloper)

	testStory := entities.NewStory("test-story", "This is a story")
	storyRepository.Create(testStory)

	changer := estimates.NewChanger(userRepository, storyRepository)
	estimateHdlr := handlers.NewEstimateHandler(router, nil, changer, nil)
	estimateHdlr.RegisterHandlers()

	estimate := 5
	payload := fmt.Sprintf(`{
			"DeveloperId": "%s",
			"Value": %d
		}`, testDeveloper.ID(), estimate)

	path := fmt.Sprintf("/v1/stories/%s/estimates", testStory.ID())
	req, err := http.NewRequest("PUT", path, strings.NewReader(payload))
	if err != nil {
		t.Error(err)
	}

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected to have HTTP Code [%d] , got [%d] ", http.StatusOK, recorder.Code)
	}
}

func TestListEstimatesHandler(t *testing.T) {
	setup()

	recorder := httptest.NewRecorder()
	router := mux.NewRouter()

	testStory := entities.NewStory("test-story", "This is a story")
	storyRepository.Create(testStory)

	lister := estimates.NewLister(storyRepository)
	estimateHdlr := handlers.NewEstimateHandler(router, nil, nil, lister)
	estimateHdlr.RegisterHandlers()

	path := fmt.Sprintf("/v1/stories/%s/estimates", testStory.ID())
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		t.Error(err)
	}

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("Expected to have HTTP Code [%d] , got [%d] ", http.StatusOK, recorder.Code)
	}
}
