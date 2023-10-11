package sqlstore

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"
)

// Test store
func TestDB(t *testing.T, databaseUrl string) (*sql.DB, func(...string)) {
	t.Helper()

	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		t.Fatal()
	}

	if err := db.Ping(); err != nil {
		t.Fatal()
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := db.Exec(fmt.Sprintf("TRUNCATE %s CASCADE ", strings.Join(tables, ", "))); err != nil {
				t.Fatal()
			}
		}

		db.Close()
	}
}
