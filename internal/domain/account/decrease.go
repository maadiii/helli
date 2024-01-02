package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/event"
)

func (s *service) Decrease(entity *entity.Account, amount float64) (err error) {
	if err = s.accounts.Decrease(entity, amount); err != nil {
		return
	}

	s.raiseDecrease(entity, amount)

	return
}

func (s *service) raiseDecrease(entity *entity.Account, amount float64) {
	go func() {
		event := event.Event{
			Action:    event.Decrease,
			Username:  entity.Username,
			AccountID: entity.ID,
			Amount:    amount,
		}

		s.event.Raise(event)
	}()
}
