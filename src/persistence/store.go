package persistence

var store Datastorer

// Datastorer interface
type Datastorer interface {
	Open(connect ...string) (Datastorer, error)
}
