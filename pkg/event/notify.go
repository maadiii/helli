package event

import "context"

func notifyEvent(ctx context.Context, event Event) {
	// TODO: send event to some message streamer like kafa, nats, rabbitmq or some other service to notify user
}
