package apiserver

import (
	"database/sql"
	"net/http"
	"net/http/pprof"

	"github.com/IRonzin/http-rest-api/internal/app/store/sqlstore"
	"github.com/gorilla/sessions"
)

// Start ...
func Start(config *Config, isNeedPprof bool) error {
	db, err := newDB(config.DatabaseUrl)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	s := newServer(store, sessionStore)
	s.configureLogger(config)

	listenError := make(chan error, 1)

	go func() {
		s.logger.Infof("starting api server on port %s", config.BindAddr)
		listenError <- http.ListenAndServe(config.BindAddr, s)
		defer close(listenError)
	}()

	go func() {
		pprofMux := http.NewServeMux()
		pprofMux.HandleFunc("/debug/pprof/", pprof.Index)
		pprofMux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		pprofMux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		pprofMux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		pprofMux.HandleFunc("/debug/pprof/trace", pprof.Trace)

		// Manually add support for paths linked to by index page at /debug/pprof/
		pprofMux.Handle("/debug/pprof/goroutine", pprof.Handler("goroutine"))
		pprofMux.Handle("/debug/pprof/heap", pprof.Handler("heap"))
		pprofMux.Handle("/debug/pprof/threadcreate", pprof.Handler("threadcreate"))
		pprofMux.Handle("/debug/pprof/block", pprof.Handler("block"))
		s.logger.Infof("starting pprof server on port %s", config.PprofBindAddr)
		if err := http.ListenAndServe(config.PprofBindAddr, pprofMux); err != nil {
			s.logger.Errorf("pprof is exited with error: %v", err)
		}
	}()

	for err := range listenError {
		return err
	}

	return nil
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
