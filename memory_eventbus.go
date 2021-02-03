package eventbus

// Memory EventBus implementation
type Memory struct {
	subscribers map[string]func(event interface{})
}

func (t *Memory) Init() {
	t.subscribers = make(map[string]func(event interface{}))
}

func (t *Memory) Publish(event *EventContainer) {
	s := t.subscribers[event.name]

	if s != nil {
		s(event)
	}
}

func (t *Memory) Subscribe(subscriber *Subscriber) {
	t.subscribers[(*subscriber).GetEventName()]=(*subscriber).AcceptEvent
}
