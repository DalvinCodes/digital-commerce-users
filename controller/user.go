package controller

import (
	"context"
	"errors"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"github.com/DalvinCodes/digital-commerce/users/service"
	"github.com/gofiber/fiber/v2"
)

type UserCtrl interface {
	Create(ctx context.Context, user *model.User) error
}

type UserController struct {
	Controller UserCtrl
}

type User struct {
	Service service.UserService
}

func NewUserController(userCtrl UserCtrl) *UserController {
	return &UserController{Controller: userCtrl}
}

func (*User) Create(ctx fiber.Ctx) error {
	return errors.New("not implemented")
}
