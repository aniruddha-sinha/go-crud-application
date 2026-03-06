package store

import "database/sql"

type UserStore struct {
	DB *sql.DB
}

type User struct {
	ID    int    `json:id`
	Name  string `json:name`
	Email string `json:email`
}
