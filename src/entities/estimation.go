package entities

// Estimation entity
type Estimation struct {
	developerID string
	storyID     string
	value       int
}

// DeveloperID for estimation
func (e Estimation) DeveloperID() string {
	return e.developerID
}

// StoryID for estimation
func (e Estimation) StoryID() string {
	return e.storyID
}

// Value of estimation
func (e Estimation) Value() int {
	return e.value
}

// NewEstimation returns a new Estimation entity
func NewEstimation(developerID, storyID string, value int) Estimation {
	return Estimation{developerID, storyID, value}
}
