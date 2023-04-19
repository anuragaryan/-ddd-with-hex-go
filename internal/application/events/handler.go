package events

//TALK: Emphasise on Observer pattern. Also Domain Events are internal Events for our Bounded Context, for external ones need infra like sns-sqs.

// EventHandler interface for describing any object that should be notified upon some Event has happened.
type EventHandler interface {
	Notify(event Event) error
}

// EventPublisher central structure for notifying all EventHandler.
type EventPublisher struct {
	handlers map[string][]EventHandler
}

// NewEventPublisher creates a new event publisher.
func NewEventPublisher() *EventPublisher {
	return &EventPublisher{
		handlers: make(map[string][]EventHandler),
	}
}

// Subscribe subscribes EventHandler to particular Events,
func (e *EventPublisher) Subscribe(handler EventHandler, events ...Event) {
	for _, event := range events {
		handlers := e.handlers[event.Name()]
		handlers = append(handlers, handler)
		e.handlers[event.Name()] = handlers
	}
}

// Notify notifies subscribed EventHandler for particular Event.
func (e *EventPublisher) Notify(event Event) error {
	for _, handler := range e.handlers[event.Name()] {
		err := handler.Notify(event)
		if err != nil {
			return err
		}
	}

	return nil
}
