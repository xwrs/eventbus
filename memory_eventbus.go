package eventbus

// Memory EventBus implementation
type Memory struct {
	subscribers map[string]func(event *IntegrationEvent)
}

func (t *Memory) Init() {
	t.subscribers = make(map[string]func(event *IntegrationEvent))
}

func (t *Memory) Publish(event *IntegrationEvent) {
	s := t.subscribers[event.Name]

	if s != nil {
		s(event)
	}
}

func (t *Memory) Subscribe(subscriber *Subscriber) {
	t.subscribers[(*subscriber).GetEventName()]=(*subscriber).AcceptEvent
}
