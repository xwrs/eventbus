package eventbus

import (
	"reflect"
)

type Subscriber interface {
	 GetEventName() string
	GetEventType() reflect.Type
	 GetTopicName() string
	 AcceptEvent(event *IntegrationEvent)
}
