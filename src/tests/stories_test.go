package tests

import (
	"entities"
	"services/stories"
	"testing"
)

func TestProductOwnerCreatesStory(t *testing.T) {
	setup()

	storyTitle := "test"

	storyCreator := stories.NewCreator(storyRepository)
	story := stories.StoryModel{
		Title:       storyTitle,
		Description: "test",
	}
	response := storyCreator.Execute(story)

	if response.Err != nil {
		t.Error(response.Err)
	}

	if response.Title != storyTitle {
		t.Errorf("Expected to have title [%s], got [%s]", storyTitle, response.Title)
	}
}

func TestListStories(t *testing.T) {
	setup()

	var testStory *entities.Story
	storyTitle := "test-story"
	storiesNo := 2

	for i := 0; i < storiesNo; i++ {
		testStory = entities.NewStory(storyTitle, "This is a story")
		storyRepository.Create(testStory)
	}

	storyLister := stories.NewLister(storyRepository)
	response := storyLister.Execute()
	if response.Err != nil {
		t.Error(response.Err)
	}

	if len(response.Titles) != storiesNo {
		t.Errorf("Expected to have [%d] stories titles, got [%d]", storiesNo, len(response.Titles))
	}
}
