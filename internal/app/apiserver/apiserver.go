package apiserver

import (
	"database/sql"
	"net/http"

	"github.com/IRonzin/http-rest-api/internal/app/store/sqlstore"
)

// Start ...
func Start(config *Config) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	s := newServer(store)
	s.configureLogger(config)

	s.logger.Info("starting api server")
	return http.ListenAndServe(config.BindAddr, s)
}

func newDB(databaseUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
