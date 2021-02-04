package eventbus


// EventBus is base implementation of IEventBus
type EventBus interface {
	Init()
	Publish(event *EventContainer)
	Subscribe(subscriber *Subscriber) error
}
