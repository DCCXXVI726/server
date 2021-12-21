package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type User struct {
	Email string
	Pass  string
	ID    int
}
type Store struct {
	db *sql.DB
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) all() (int, string) {
	var (
		id    int
		email string
	)
	s.db.QueryRow("select id, email from users LIMIT 1").Scan(&id, &email)
	return id, email
}

func (s *Store) Open() error {
	databaseUrl := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", databaseUrl)
	if err != nil {
		return fmt.Errorf("can't open db: %s", err)
	}

	if err := db.Ping(); err != nil {
		return fmt.Errorf("can't ping db: %s", err)
	}

	s.db = db
	return nil
}

func (s *Store) Close() {
	s.db.Close()
}
