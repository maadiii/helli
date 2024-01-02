package account_test

import (
	"testing"

	"github.com/maadiii/helli/internal/entity"
)

func TestTransfer(t *testing.T) {
	firstEntity := &entity.Account{Username: "user1", Fund: 100}
	if err := dal.Insert(firstEntity); err != nil {
		t.Error(err)
	}

	secondEntity := &entity.Account{Username: "user2", Fund: 100}
	if err := dal.Insert(secondEntity); err != nil {
		t.Error(err)
	}

	t.Run("first transfer", func(t *testing.T) {
		t.Parallel()

		if err := dal.Transfer(firstEntity, secondEntity, 30); err != nil {
			t.Error(err)
		}
	})

	t.Run("second transfer", func(t *testing.T) {
		t.Parallel()

		if err := dal.Transfer(firstEntity, secondEntity, 20); err != nil {
			t.Error(err)
		}
	})

	t.Run("transfers result", func(t *testing.T) {
		t.Parallel()

		first := &entity.Account{Username: "user1", ID: 1}
		second := &entity.Account{Username: "user2", ID: 1}

		if err := dal.Get(second); err != nil {
			t.Error(err)
		}

		if err := dal.Get(first); err != nil {
			t.Error(err)
		}

		if first.Fund != 50 && first.Fund != 70 {
			t.Errorf("expect first account fund be %d or %d, got %f", 50, 70, first.Fund)
		}

		if second.Fund != 150 && second.Fund != 130 {
			t.Errorf("expect second account fund %d or %d, got %f", 150, 130, second.Fund)
		}
	})
}
