package user_test

import (
	"testing"

	"github.com/maadiii/helli/internal/entity"
	"github.com/maadiii/hertzwrapper/errors"
)

func TestInsert(t *testing.T) {
	t.Run("insert first user", func(t *testing.T) {
		t.Parallel()

		if err := dal.Insert(&entity.User{Username: "first_user"}); err != nil {
			t.Error(err)
		}
	})

	t.Run("insert second user", func(t *testing.T) {
		t.Parallel()

		if err := dal.Insert(&entity.User{Username: "second_user"}); err != nil {
			t.Error(err)
		}
	})

	t.Run("insert third user", func(t *testing.T) {
		t.Parallel()

		if err := dal.Insert(&entity.User{Username: "third_user"}); err != nil {
			t.Error(err)
		}
	})

	t.Run("insert fourth user", func(t *testing.T) {
		t.Parallel()

		if err := dal.Insert(&entity.User{Username: "fourth_user"}); err != nil {
			t.Error(err)
		}
	})

	t.Run("insert repetitious user", func(t *testing.T) {
		t.Parallel()

		if err := dal.Insert(&entity.User{Username: "first_user"}); !errors.Is(err, errors.Conflict) {
			t.Error("expected conflic error got", err.Error())
		}
	})
}
