package service

import (
	"context"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"github.com/DalvinCodes/digital-commerce/users/repo"
)

type UserServiceI interface {
	Create(ctx context.Context, user *model.User) error
	ListAll(ctx context.Context) ([]*model.User, error)
}

type UserService struct {
	Repo repo.UserRepository
}

func NewUser(repo repo.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) Create(ctx context.Context, user *model.User) error {
	return s.Repo.Create(ctx, user)
}

func (s *UserService) ListAll(ctx context.Context) ([]*model.User, error) {
	return s.Repo.ListAll(ctx)
}
