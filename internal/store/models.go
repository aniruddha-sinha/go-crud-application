package store

import "database/sql"

type UserStore struct {
	DB *sql.DB
}

type User struct {
	ID    int
	Name  string
	Email string
}
