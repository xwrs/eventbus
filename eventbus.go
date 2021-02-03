package eventbus


// EventBus is base implementation of IEventBus
type EventBus interface {
	Init()
	Publish(event *IntegrationEvent)
	Subscribe(subscriber *Subscriber)
}
