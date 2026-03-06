package database

import (
	"context"
	"database/sql"
	"time"
)

func InitDB(dsn string) (*sql.DB, func()) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		panic(err)
	}

	//ping to ensure dsn is correct
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		panic("Could not connect to Postgres: " + err.Error())
	}

	return db, func() { db.Close() }
}
