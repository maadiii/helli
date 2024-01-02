package event

import (
	"context"
)

type Raiser interface {
	Raise(event Event)
}

func NewRaiser(ctx context.Context, worker int) Raiser {
	e := &raiser{
		channel: make(chan Event),
	}

	for i := 0; i < worker; i++ {
		e.run(ctx)
	}

	go func() {
		<-ctx.Done()

		close(e.channel)
	}()

	return e
}

type raiser struct {
	channel chan Event
}

func (e *raiser) run(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case event := <-e.channel:
				logEvent(ctx, event)
				notifyEvent(ctx, event)
			}
		}
	}()
}

func (e *raiser) Raise(event Event) {
	e.channel <- event
}

type action string

const (
	Create   action = "create"
	Increase action = "increase"
	Decrease action = "decrease"
	Transfer action = "transfer"
	receive  action = "receive"
)

type Event struct {
	Username             string  `json:"username,omitempty"`
	SourceUsername       string  `json:"source_username,omitempty"`
	DestinationUsername  string  `json:"distination_username,omitempty"`
	AccountID            int     `json:"account_id,omitempty"`
	SourceAccountID      int     `json:"source_account_id,omitempty"`
	DestinationAccountID int     `json:"distination_account_id,omitempty"`
	Action               action  `json:"action,omitempty"`
	Amount               float64 `json:"amount,omitempty"`
}
