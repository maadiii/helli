package account_test

import (
	"os"
	"testing"

	"github.com/maadiii/helli/internal/data_access/account"
)

var dal account.DataAccess

func TestMain(m *testing.M) {
	dal = account.New()

	os.Exit(m.Run())
}
