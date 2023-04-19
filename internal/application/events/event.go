package events

// Event interface for describing Domain Event
type Event interface {
	Name() string
}
