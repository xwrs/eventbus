package eventbus

import "reflect"

type Subscriber interface {
	 GetEventName() string
	GetEventType() reflect.Type
	 GetTopicName() string
	 AcceptEvent(event *IntegrationEvent)
}

type DummyEventSubscriber struct {
	eventType reflect.Type
	Subscriber
}

func (*DummyEventSubscriber) GetEventName() string {
	event := DummyPayload{}
	return reflect.TypeOf(event).Name()
}

func (*DummyEventSubscriber) GetTopicName() string {
	event := DummyPayload{}
	return reflect.TypeOf(event).PkgPath()
}

func (s *DummyEventSubscriber) AcceptEvent(event *IntegrationEvent)  {
	return
}

func NewDummyEventSubscriber() (s *DummyEventSubscriber) {
	event := DummyPayload{}
	return &DummyEventSubscriber{
		eventType: reflect.TypeOf(event),
	}
}