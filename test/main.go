package main

import (
	"github.com/xwrs/eventbus"
	"time"
)

func main() {
	var eventBus eventbus.EventBus

	eventBus = &eventbus.AzureServiceBus{}
	eventBus.Init()

	var sub eventbus.Subscriber
	sub = NewDummyEventSubscriber()

	eventBus.Subscribe(&sub)
	eventBus.Publish(eventbus.NewEventContainer(&DummyEvent{
		A: "test a",
		B: "test b",
	}, nil))
	time.Sleep(time.Second*30)

}
