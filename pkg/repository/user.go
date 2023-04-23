package repository

import (
	"context"
	"echo-template/pkg/domain"

	"github.com/jackc/pgx/v4/pgxpool"
)

type Repository interface {
	CreateUser(ctx context.Context, request domain.UserRequest) (domain.User, error)
	GetUser(ctx context.Context, request int) (domain.User, error)
}

type repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, request domain.UserRequest) (domain.User, error) {
	var response domain.User

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING *`
	if err := r.db.QueryRow(ctx, query, request.Name, request.Email, request.Password).Scan(&response); err != nil {
		return domain.User{}, err
	}

	return response, nil
}
