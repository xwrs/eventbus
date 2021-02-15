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

func (t *AzureServiceBus) Publish(event *EventContainer) {
	ctx := context.Background()
	eventName := event.GetName()
	topicName := event.GetTopicName()

	topic, _ := t.client.NewTopic(topicName)
	messageBytes, _ := json.Marshal(event.Payload)
	message := servicebus.NewMessage(messageBytes)
	message.Label = eventName
	if event.Metadata != nil {
		for k, v := range *event.Metadata {
			message.UserProperties[k] = v
		}
	}
	topic.Send(ctx, message)
}

func (t *AzureServiceBus) Subscribe(subscriber *Subscriber) error {
	ctx := context.Background()
	eventName := (*subscriber).GetEventName()
	eventType := (*subscriber).GetEventType()
	topicName := (*subscriber).GetTopicName()

	t.client.NewTopicManager().Put(ctx, topicName)
	sm, err := t.client.NewSubscriptionManager(topicName)
	if err != nil {
		return err
	}

	_, err = sm.Get(ctx, eventName)
	if err != nil {
		if _, ok := err.(servicebus.ErrNotFound); ok {
			_, err := sm.Put(ctx, eventName, func(description *servicebus.SubscriptionDescription) error {
				ruleName := "EventName"
				ruleDescription := &servicebus.DefaultRuleDescription{
					Filter: servicebus.CorrelationFilter{
						Label: &eventName,
					}.ToFilterDescription(),
					Name: &ruleName,
				}
				description.DefaultRuleDescription = ruleDescription
				return nil
			})

			if err != nil {
				return err
			}
		}
	}

	topic, err := t.client.NewTopic(topicName)
	if err != nil {
		return err
	}
	subscription, err := topic.NewSubscription(eventName)
	if err != nil {
		return err
	}

	receiver, err := subscription.NewReceiver(ctx)
	if err != nil {
		return err
	}
	var handlerFunc servicebus.HandlerFunc
	handlerFunc = func(ctx context.Context, message *servicebus.Message) error {
		payload := reflect.New(eventType).Interface()
		if err := json.Unmarshal(message.Data, payload); err != nil {
			panic(err)
		}
		message.Complete(ctx)
		(*subscriber).AcceptEvent(payload)
		return nil
	}
	receiver.Listen(ctx, handlerFunc)
	return nil
}
