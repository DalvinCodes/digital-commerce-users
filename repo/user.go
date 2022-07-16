package repo

import (
	"context"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	ListAll(ctx context.Context) ([]*model.User, error)
	FindByID(ctx context.Context, id string) (*model.User, error)
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
}

type UserRepo struct {
	DB *gorm.DB
}

const (
	idIs       = `id = ?`
	usernameIs = `username = ?`
)

func NewUserRepository(db *gorm.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	return r.DB.Debug().Create(&user).WithContext(ctx).Error
}

func (r *UserRepo) ListAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	err := r.DB.Debug().
		Find(&users).
		WithContext(ctx).
		Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepo) FindByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	if err := r.DB.Debug().WithContext(ctx).
		Where(idIs, id).
		Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	var user model.User

	if err := r.DB.Debug().WithContext(ctx).
		Where(usernameIs, username).
		Find(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) Update(ctx context.Context, user *model.User) error {
	return r.DB.Debug().WithContext(ctx).Model(&user).Updates(&user).Error
}

func (r *UserRepo) Delete(ctx context.Context, user *model.User) error {
	return r.DB.Debug().WithContext(ctx).Delete(&user).Error
}
