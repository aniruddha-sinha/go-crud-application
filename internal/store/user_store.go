package store

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func NewUserStore(dsn string) (*UserStore, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	return &UserStore{DB: db}, nil //* update the UserStore value
}
