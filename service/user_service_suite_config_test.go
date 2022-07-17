package service

import (
	"github.com/DalvinCodes/digital-commerce/users/model"
	"github.com/DalvinCodes/digital-commerce/users/repo/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserServiceTestSuite struct {
	suite.Suite
	Repo    mocks.UserRepository
	Service UserService
	User    model.User
}

func (s *UserServiceTestSuite) SetupSuite() {
	s.Repo = mocks.UserRepository{}
	s.Service = UserService{Repo: &s.Repo}
}

func TestRunUserServiceTestSuite(t *testing.T) {
	suite.Run(t, &UserServiceTestSuite{})
}

func (s *UserServiceTestSuite) SeedMockUserData() *model.User {
	user := s.User
	user.ID = gofakeit.UUID()
	user.Username = gofakeit.Username()
	user.FirstName = gofakeit.FirstName()
	user.LastName = gofakeit.LastName()
	user.Email = gofakeit.Email()
	user.DateOfBirth = gofakeit.Date().Format("01/02/2006")

	return &user
}
