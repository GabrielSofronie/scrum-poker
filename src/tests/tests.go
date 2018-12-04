package tests

import (
	"persistence"
	"persistence/inmemory"
	"repositories"
)

var store persistence.Datastorer
var storyRepository repositories.StoryRepository
var userRepository repositories.UserRepository

func setup() {
	store = &inmemory.InMemory{}
	storyRepository = inmemory.NewStories(store)
	userRepository = inmemory.NewUsers(store)
}
