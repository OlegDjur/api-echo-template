package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func NewPsqlDB() *pgxpool.Pool {
	connStr := "postgres://olegdjur:qwerty@localhost:5432/users?sslmode=disable"

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err)
	}
	// defer db.Close()

	return db
}
