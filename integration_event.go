package eventbus

import "reflect"

type IntegrationEvent struct {
	name      string
	topicName string
	Metadata  *map[string]interface{}
	Payload   interface{}
}

func (e *IntegrationEvent) GetName() string {
	return e.name
}

func (e *IntegrationEvent) GetTopicName() string {
	return e.topicName
}

func NewIntegrationEvent(payload interface{}, metadata *map[string]interface{}) *IntegrationEvent {
	return &IntegrationEvent{
		name:      reflect.TypeOf(payload).Name(),
		topicName: reflect.TypeOf(payload).PkgPath(),
		Metadata:  metadata,
		Payload:   payload,
	}
}
