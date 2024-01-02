package account

import (
	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/helli/pkg/event"
)

func (s *service) Transfer(src *entity.Account, dstUserID, dstAccountID int, amount float64) (err error) {
	user := &entity.User{ID: dstUserID}
	if err = s.users.GetByID(user); err != nil {
		return
	}

	dst := &entity.Account{ID: dstAccountID, Username: user.Username}

	if err = s.accounts.Transfer(src, dst, amount); err != nil {
		return
	}

	s.raiseTransfer(src, dst, amount)

	return
}

func (s *service) raiseTransfer(src *entity.Account, dst *entity.Account, amount float64) {
	go func() {
		event := event.Event{
			Action:               event.Transfer,
			SourceUsername:       src.Username,
			SourceAccountID:      src.ID,
			DestinationUsername:  dst.Username,
			DestinationAccountID: dst.ID,
			Amount:               amount,
		}

		s.event.Raise(event)
	}()
}
