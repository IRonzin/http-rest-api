package sqlstore_test

import (
	"os"
	"testing"
)

var (
	databaseUrl string
)

func TestMain(m *testing.M) {
	databaseUrl = os.Getenv("TEST_DATABASE_URL")
	if databaseUrl == "" {
		databaseUrl = "host=localhost port=5431 user=admin_test password=12345 dbname=restapi_test sslmode=disable"
	}

	os.Exit(m.Run())
}
