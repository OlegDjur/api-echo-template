package repository

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"honnef.co/go/tools/config"
)

func NewPsqlDB(c config.Config) *pgxpool.Pool {
	connStr := "postgres://olegdjur:qwerty@localhost:5432/postgres?sslmode=disable"

	db, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database:", err)
	}
	defer db.Close()

	return db
}
