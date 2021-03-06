package service

import (
	"context"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"github.com/DalvinCodes/digital-commerce/users/repo"
)

type UserServiceI interface {
	Create(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	ListAll(ctx context.Context) ([]*model.User, error)
	Update(ctx context.Context, user *model.User) error
	UpdateUsername(ctx context.Context, userID, username string) error
	UpdateEmail(ctx context.Context, userID, email string) error
	UpdatePassword(ctx context.Context, userID, password string) error
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

func (s *UserService) FindByID(ctx context.Context, id string) (*model.User, error) {
	return s.Repo.FindByID(ctx, id)
}

func (s *UserService) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	return s.Repo.FindByUsername(ctx, username)
}

func (s *UserService) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.Repo.FindByEmail(ctx, email)
}

func (s *UserService) Delete(ctx context.Context, user *model.User) error {
	return s.Repo.Delete(ctx, user)
}

func (s *UserService) Update(ctx context.Context, user *model.User) error {
	return s.Repo.Update(ctx, user)
}

func (s *UserService) UpdateUsername(ctx context.Context, userID, username string) error {
	user := &model.User{}

	user.ID = userID
	user.Username = username

	return s.Repo.Update(ctx, user)
}

func (s *UserService) UpdateEmail(ctx context.Context, userID, email string) error {
	user := &model.User{}

	user.ID = userID
	user.Email = email

	return s.Repo.Update(ctx, user)
}

func (s *UserService) UpdatePassword(ctx context.Context, userID, password string) error {
	user := &model.User{}

	user.ID = userID
	user.Password = password

	return s.Repo.Update(ctx, user)
}
