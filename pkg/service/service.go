package service

import (
	"echo-template/pkg/domain"
	"echo-template/pkg/repository"

	"golang.org/x/net/context"
)

type Service interface {
	CreateUser(ctx context.Context, request domain.UserRequest) (domain.User, error)
	GetUser(ctx context.Context, request int) (domain.User, error)
}

type service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *service {
	return &service{repo: repo}
}

func (s *service) CreateUser(ctx context.Context, request domain.UserRequest) (domain.User, error) {
	return s.repo.CreateUser(ctx, request)
}

func (s *service) GetUser(ctx context.Context, request int) (domain.User, error) {
	return s.repo.GetUser(ctx, request)
}
