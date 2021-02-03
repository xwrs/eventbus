package main

import (
	"github.com/xwrs/eventbus"
	"reflect"
)

type DummyEventSubscriber struct {
	eventType reflect.Type
	eventbus.Subscriber
}

func (*DummyEventSubscriber) GetEventName() string {
	event := &DummyPayload{}
	return reflect.TypeOf(event).Elem().Name()
}

func (*DummyEventSubscriber) GetTopicName() string {
	event := &DummyPayload{}
	return reflect.TypeOf(event).Elem().PkgPath()
}

func(*DummyEventSubscriber) GetEventType() reflect.Type  {
	event := &DummyPayload{}
	return reflect.TypeOf(event)
}

func (s *DummyEventSubscriber) AcceptEvent(event *eventbus.IntegrationEvent)  {
	return
}

func NewDummyEventSubscriber() (s *DummyEventSubscriber) {
	event := &DummyPayload{}
	return &DummyEventSubscriber{
		eventType: reflect.TypeOf(event),
	}
}
