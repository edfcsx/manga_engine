package events

import "github.com/google/uuid"

type SubscriptionAction interface {
	Handler(params interface{})
}

type Subscription struct {
	id     string
	action SubscriptionAction
}

type Unsubscription struct {
	event string
	id    string
}

var handlers = map[string][]Subscription{}

func Subscribe(event string, action SubscriptionAction) Unsubscription {
	if handlers[event] == nil {
		handlers[event] = make([]Subscription, 10)
	}

	id := uuid.New().String()
	subscription := Subscription{
		id:     id,
		action: action,
	}

	handlers[event] = append(handlers[event], subscription)

	return Unsubscription{
		id:    id,
		event: event,
	}
}

func Unsubscribe(unsubscribe Unsubscription) {
	if subscriptions, ok := handlers[unsubscribe.event]; ok {
		for i, subscription := range subscriptions {
			if subscription.id == unsubscribe.id {
				handlers[unsubscribe.event] = append(subscriptions[:i], subscriptions[i+1:]...)
				break
			}
		}
	}
}

func UnsubscribeAll() {
	handlers = map[string][]Subscription{}
}
