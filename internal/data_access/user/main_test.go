package user_test

import (
	"os"
	"testing"

	"github.com/maadiii/helli/internal/data_access/user"
)

var dal user.DataAccess

func TestMain(m *testing.M) {
	dal = user.New()

	os.Exit(m.Run())
}
