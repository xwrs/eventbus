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

func NewEventContainerReflected(payload interface{}, metadata *map[string]interface{}) *EventContainer {
	eventName := reflect.TypeOf(payload).Elem().Name()
	topicName := reflect.TypeOf(payload).Elem().PkgPath()
	return NewEventContainer(eventName, topicName, payload, metadata)
}

func NewEventContainer(name string, topicName string, payload interface{}, metadata *map[string]interface{}) *EventContainer {
	return &EventContainer{
		name:      name,
		topicName: topicName,
		Metadata:  metadata,
		Payload:   payload,
	}
}
