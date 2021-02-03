package eventbus

import (
	"context"
	"encoding/json"
	"reflect"

	servicebus "github.com/Azure/azure-service-bus-go"
)

type AzureServiceBus struct {
	client *servicebus.Namespace
}

func (t *AzureServiceBus) Init() {
	client, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(""))
	if err != nil {
		panic("fail")
	}
	t.client = client
}

func (t *AzureServiceBus) Publish(event *IntegrationEvent) {
	ctx := context.Background()
	eventName := event.GetName()
	topicName := event.GetTopicName()

	topic, _ := t.client.NewTopic(topicName)
	messageBytes, _ := json.Marshal(event)
	message := servicebus.NewMessage(messageBytes)
	message.Label = eventName
	if *event.Metadata != nil {
		for k, v := range *event.Metadata {
			message.UserProperties[k] = v
		}
	}
	topic.Send(ctx, message)
}

func (t *AzureServiceBus) Subscribe(subscriber *Subscriber) {
	ctx := context.Background()
	eventName := (*subscriber).GetEventName()
	eventType := (*subscriber).GetEventType()
	topicName := (*subscriber).GetTopicName()

	t.client.NewTopicManager().Put(ctx, topicName)
	if sm, err := t.client.NewSubscriptionManager(topicName); err == nil {
		if _, err := sm.Get(ctx, eventName); err != nil {
			if _, ok := err.(servicebus.ErrNotFound); ok {
				if _, err := sm.Put(ctx, eventName, func(description *servicebus.SubscriptionDescription) error {
					ruleName := "EventName"
					ruleDescription := &servicebus.DefaultRuleDescription{
						Filter: servicebus.CorrelationFilter{
							Label: &eventName,
						}.ToFilterDescription(),
						Name: &ruleName,
					}
					description.DefaultRuleDescription = ruleDescription
					return nil
				}); err != nil {
					panic(err)
				}
			}
		}
	} else {
		panic(err)
	}

	if topic, err := t.client.NewTopic(topicName); err == nil {
		if subscription, err := topic.NewSubscription(eventName); err == nil {
			if receiver, err := subscription.NewReceiver(ctx); err == nil {
				var handlerFunc servicebus.HandlerFunc
				handlerFunc = func(ctx context.Context, message *servicebus.Message) error {
					payload := reflect.New(eventType)
					json.Unmarshal(message.Data, payload)
					event := NewIntegrationEvent(payload, &message.UserProperties)

					message.Complete(ctx)
					(*subscriber).AcceptEvent(event)
					return nil
				}
				receiver.Listen(ctx, handlerFunc)
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
}
