package entities

import "workmate"

// Discussion entity
type Discussion struct {
	id        string
	userID    string
	storyID   string
	replyToID string
	content   string
}

// ID of discussion
func (d Discussion) ID() string {
	return d.id
}

// UserID creating the discussion
func (d Discussion) UserID() string {
	return d.userID
}

// StoryID of discussion
func (d Discussion) StoryID() string {
	return d.storyID
}

// ReplyID to discussion
func (d Discussion) ReplyID() string {
	return d.replyToID
}

// Content of discussion
func (d Discussion) Content() string {
	return d.content
}

// NewQuestion returns a Discussion interface
func NewQuestion(userID, storyID, content string) Discussion {
	return Discussion{
		id:      workmate.SimpleID(),
		userID:  userID,
		storyID: storyID,
		content: content,
	}
}

// NewAnswer returns a Discussion interface
func NewAnswer(userID, storyID, replyID, content string) Discussion {
	return Discussion{
		id:        workmate.SimpleID(),
		userID:    userID,
		storyID:   storyID,
		replyToID: replyID,
		content:   content,
	}
}
