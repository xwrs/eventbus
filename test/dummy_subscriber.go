package main

import (
	"fmt"
	"github.com/xwrs/eventbus"
	"reflect"
)

type DummyEventSubscriber struct {
	eventType reflect.Type
	eventbus.Subscriber
}

func (*DummyEventSubscriber) GetEventName() string {
	event := &DummyEvent{}
	return reflect.TypeOf(event).Elem().Name()
}

func (*DummyEventSubscriber) GetTopicName() string {
	event := &DummyEvent{}
	return reflect.TypeOf(event).Elem().PkgPath()
}

func(*DummyEventSubscriber) GetEventType() reflect.Type  {
	event := &DummyEvent{}
	return reflect.TypeOf(event).Elem()
}

func (s *DummyEventSubscriber) AcceptEvent(event interface{})  {
	dummyEvent := event.(*DummyEvent)
	fmt.Print(dummyEvent)
	return
}

func NewDummyEventSubscriber() (s *DummyEventSubscriber) {
	event := &DummyEvent{}
	return &DummyEventSubscriber{
		eventType: reflect.TypeOf(event),
	}
}
