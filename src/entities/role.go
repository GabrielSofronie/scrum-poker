package entities

// Role type
type Role int

// User roles
const (
	Developer Role = iota
	Moderator
	Owner
)
