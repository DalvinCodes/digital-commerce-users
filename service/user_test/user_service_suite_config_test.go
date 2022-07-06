package service

import (
	"github.com/DalvinCodes/digital-commerce/users/repo/mocks"
	"github.com/stretchr/testify/suite"
	"testing"
)

type UserServiceTestSuite struct {
	suite.Suite
	repo mocks.UserRepository
}

func (s *UserServiceTestSuite) SetupSuite() {
}

func TestRunUserServiceTestSuite(t *testing.T) {
	suite.Run(t, &UserServiceTestSuite{})
}
