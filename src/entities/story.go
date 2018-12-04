package entities

import (
	"workmate"
)

// Story entity
type Story struct {
	id          string
	title       string
	description string
	estimates   []*Estimation
	discussions []*Discussion
}

// ID of story
func (s Story) ID() string {
	return s.id
}

// Title of story
func (s Story) Title() string {
	return s.title
}

// Description of story
func (s Story) Description() string {
	return s.description
}

// Estimates of story
func (s Story) Estimates() []*Estimation {
	return s.estimates
}

// SetEstimate for story
func (s *Story) SetEstimate(estimate Estimation) {
	s.estimates = append(s.estimates, &estimate)
}

// DeleteEstimate from story
func (s *Story) DeleteEstimate(estimate Estimation) {
	for i, e := range s.estimates {
		if e.DeveloperID() == estimate.DeveloperID() {
			s.estimates[i] = s.estimates[len(s.estimates)-1]
			s.estimates[len(s.estimates)-1] = nil
			s.estimates = s.estimates[:len(s.estimates)-1]
		}
	}
}

// Discussions for a story
func (s Story) Discussions() []*Discussion {
	return s.discussions
}

// AddDiscussion to story
func (s *Story) AddDiscussion(d Discussion) {
	s.discussions = append(s.discussions, &d)
}

// DeleteDiscussion from story
func (s *Story) DeleteDiscussion(d Discussion) {
	for i, disq := range s.discussions {
		if disq.ID() == d.ID() {
			s.discussions[i] = s.discussions[len(s.discussions)-1]
			s.discussions[len(s.discussions)-1] = nil
			s.discussions = s.discussions[:len(s.discussions)-1]
		}
	}
}

// NewStory creates a new user story entity
func NewStory(title, description string) *Story {
	return &Story{
		id:          workmate.SimpleID(),
		title:       title,
		description: description,
	}
}
