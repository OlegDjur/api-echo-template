package repository

import (
	"context"
	"echo-template/pkg/domain"
	"errors"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v4/pgxpool"
)

var ErrUserNotFound = errors.New("error: User not found")

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
	if err := r.db.QueryRow(ctx, query, request.Name, request.Email, request.Password).Scan(&response.ID, &response.Name, &response.Email, &response.Password); err != nil {
		return domain.User{}, err
	}

	return response, nil
}

func (r *repository) GetUser(ctx context.Context, id int) (domain.User, error) {
	var response domain.User

	if err := r.db.QueryRow(ctx, "SELECT * FROM users WHERE id = $1", id).Scan(&response.ID, &response.Name, &response.Email, &response.Password); err != nil {
		if err == pgx.ErrNoRows {
			return domain.User{}, ErrUserNotFound
		}
		return domain.User{}, err
	}

	return response, nil
}
