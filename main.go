package eventbus

import (
	"time"
)

func main() {
	var eventBus EventBus

	eventBus = &Memory{
	}

	var sub Subscriber
	sub = NewDummyEventSubscriber()

	eventBus.Subscribe(&sub)
	eventBus.Publish(NewIntegrationEvent(&DummyPayload{
		A: "test a",
		B: "test b",
	}, nil))
	time.Sleep(time.Second*30)

}
