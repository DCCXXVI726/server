package apiserver

import (
	"database/sql"
	"net/http"
	"os"
	"github.com/gorilla/sessions"
	"github.com/sleonia/Matcha/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	defer db.Close()

	store := sqlstore.New(db)
	sessionStore := sessions.NewCookieStore([]byte(config.SessionKey))
	srv := newServer(store, sessionStore)

	bindAddr :=  os.Getenv("PORT")
	if bindAddr == "" {
		bindAddr = config.BindAddr
	}
	return http.ListenAndServe(":" + bindAddr, srv)
}

func newDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
