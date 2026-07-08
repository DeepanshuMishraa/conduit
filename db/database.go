package db

import (
	"context"
	"database/sql"
	_ "github.com/lib/pq"
	"time"
)

func ConnectDB(dbUrl string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbUrl)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)

	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
