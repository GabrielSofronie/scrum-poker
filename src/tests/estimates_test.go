package tests

import (
	"entities"
	"services/estimates"
	"testing"
)

func TestDeveloperEstimatesStory(t *testing.T) {
	setup()

	// Add test developer and story
	testDeveloper := entities.NewDeveloper("test-dev")
	userRepository.Create(testDeveloper)

	testStory := entities.NewStory("test-story", "This is a story")
	storyRepository.Create(testStory)

	// Estimate
	model := estimates.EstimateModel{
		DeveloperID: testDeveloper.ID(),
		StoryID:     testStory.ID(),
		Value:       1,
	}
	estimator := estimates.NewEstimator(userRepository, storyRepository)
	response := estimator.Execute(model)

	if response.Err != nil {
		t.Errorf(response.Err.Error())
	}

	if response.Developer != testDeveloper.Name() {
		t.Errorf("Expected to have developer name: [%s] , got [%s] ", testDeveloper.Name(), response.Developer)
	}

	if response.Story != testStory.Title() {
		t.Errorf("Expected to have story title: [%s] , got [%s] ", testStory.Title(), response.Story)
	}

	if response.Estimate != 1 {
		t.Errorf("Expected to have estimate: [1] , got [%d] ", response.Estimate)
	}
}

func TestDeveloperChangesEstimate(t *testing.T) {
	setup()

	// Add test developer and story
	testDeveloper := entities.NewDeveloper("test-dev")
	userRepository.Create(testDeveloper)

	testStory := entities.NewStory("test-story", "This is a story")
	testEstimate := entities.NewEstimation(testDeveloper.ID(), testStory.ID(), 5)
	testStory.SetEstimate(testEstimate)
	storyRepository.Create(testStory)

	// Estimate
	model := estimates.ChangeModel{
		DeveloperID: testDeveloper.ID(),
		StoryID:     testStory.ID(),
		Value:       1,
	}
	changer := estimates.NewChanger(userRepository, storyRepository)
	response := changer.Execute(model)

	if response.Err != nil {
		t.Errorf(response.Err.Error())
	}

	if response.Developer != testDeveloper.Name() {
		t.Errorf("Expected to have developer name: [%s] , got [%s] ", testDeveloper.Name(), response.Developer)
	}

	if response.Story != testStory.Title() {
		t.Errorf("Expected to have story title: [%s] , got [%s] ", testStory.Title(), response.Story)
	}

	if response.Estimate != 1 {
		t.Errorf("Expected to have estimate: [1] , got [%d] ", response.Estimate)
	}
}

func TestListEstimates(t *testing.T) {
	setup()

	// Add test developers and story
	testDevOne := entities.NewDeveloper("test-dev-one")
	testDevTwo := entities.NewDeveloper("test-dev-two")

	testStory := entities.NewStory("test-story", "This is a story")
	estimateOne := 5
	testEstimateOne := entities.NewEstimation(testDevOne.ID(), testStory.ID(), estimateOne)
	testStory.SetEstimate(testEstimateOne)
	estimateTwo := 3
	testEstimateTwo := entities.NewEstimation(testDevTwo.ID(), testStory.ID(), estimateTwo)
	testStory.SetEstimate(testEstimateTwo)
	storyRepository.Create(testStory)

	// Estimate
	model := estimates.ListModel{
		StoryID: testStory.ID(),
	}
	lister := estimates.NewLister(storyRepository)
	response := lister.Execute(model)

	if response.Err != nil {
		t.Errorf(response.Err.Error())
	}

	if len(response.Estimates) != 2 {
		t.Errorf("Expected to have 2 estimates, got [%d]", len(response.Estimates))
	}

	if response.Estimates[testDevOne.ID()] != estimateOne {
		t.Errorf("Expected to have first dev estimate: [%d], got [%d]", estimateOne, response.Estimates[testDevOne.ID()])
	}

	if response.Estimates[testDevTwo.ID()] != estimateTwo {
		t.Errorf("Expected to have second dev estimate: [%d], got [%d]", estimateTwo, response.Estimates[testDevTwo.ID()])
	}
}
