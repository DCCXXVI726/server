package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("URI")
	if databaseURL == "" {
		databaseURL = "postgres://thorker:newPassword@localhost/restapi_test1?sslmode=disable"
	}

	os.Exit(m.Run())
}
