package repository

import "github.com/jackc/pgx/v4/pgxpool"

type Repository interface {
}

type repository struct {
	db *pgxpool.Pool
}
