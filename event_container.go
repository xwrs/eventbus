package eventbus

import "reflect"

type EventContainer struct {
	name      string
	topicName string
	Metadata  *map[string]interface{}
	Payload   interface{}
}

func (e *EventContainer) GetName() string {
	return e.name
}

func (e *EventContainer) GetTopicName() string {
	return e.topicName
}

func NewEventContainer(payload interface{}, metadata *map[string]interface{}) *EventContainer {
	return &EventContainer{
		name:      reflect.TypeOf(payload).Elem().Name(),
		topicName: reflect.TypeOf(payload).Elem().PkgPath(),
		Metadata:  metadata,
		Payload:   payload,
	}
}
