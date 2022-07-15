package service

import (
	"context"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"github.com/stretchr/testify/mock"
)

func (s *UserServiceTestSuite) TestCreate_Constructor() {
	// Given
	want := &UserService{Repo: &s.Repo}

	// When
	got := NewUser(&s.Repo)

	// Then
	s.Require().Equal(want, got)
}

func (s *UserServiceTestSuite) TestCreate_SuccessfulCreate() {
	// Given
	mockUser := s.SeedMockUserData()

	// When
	s.Repo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).
		Return(nil).Once()
	err := s.Service.Create(context.Background(), mockUser)

	// Then
	s.Require().NoError(err)
}

func (s *UserServiceTestSuite) TestListAll() {
	// Given
	var want []*model.User

	for i := 0; i < 5; i++ {
		want = append(want, s.SeedMockUserData())
	}

	// When
	s.Repo.On("ListAll", mock.AnythingOfType("*context.emptyCtx")).
		Return(want, nil).Once()

	users, err := s.Service.ListAll(context.Background())

	// Then
	s.Require().NoError(err)
	s.Require().Equal(len(want), len(users))
}
