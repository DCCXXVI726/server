package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

type Store struct {
	db *sql.DB
}

func NewStore() *Store {
	return &Store{}
}

func (s *Store) FindUserByEmail(email string) (User, error) {
	var user User
	if err := s.db.QueryRow(
		"select id, email, encrypted_password from users where email = $1",
		email,
	).Scan(
		&user.Id,
		&user.Email,
		&user.EncryptedPassword,
	); err != nil {
		return User{}, fmt.Errorf("can't find user with email %s : %s", email, err)
	}
	return user, nil
}

func (s *Store) CreateUser(user User) error {

	return nil
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
