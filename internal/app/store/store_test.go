package store_test

import (
	"os"
	"testing"
)

var (
	databaseURL string
)

func TestMain(m *testing.M) {
	databaseURL = os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// dbname should be restapi_test
		databaseURL = "host=172.28.1.5 user=admin password=123 dbname=restapi_dev sslmode=disable"
	}

	os.Exit(m.Run())
}
