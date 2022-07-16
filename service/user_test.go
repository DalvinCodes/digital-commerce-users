package service

import (
	"context"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"github.com/stretchr/testify/mock"
)

func (s *UserServiceTestSuite) TestUserService_CreateConstructor() {
	// Given
	want := &UserService{Repo: &s.Repo}

	// When
	got := NewUser(&s.Repo)

	// Then
	s.Require().Equal(want, got)
}

func (s *UserServiceTestSuite) TestUserService_Create_Successful() {
	// Given
	mockUser := s.SeedMockUserData()

	// When
	s.Repo.On("Create", mock.AnythingOfType("*context.emptyCtx"), mockUser).
		Return(nil).Once()
	err := s.Service.Create(context.Background(), mockUser)

	// Then
	s.Require().NoError(err)
}

func (s *UserServiceTestSuite) TestUserService_ListAll() {
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

func (s *UserServiceTestSuite) TestUserService_FindByID() {
	// Given
	want := s.SeedMockUserData()

	// When
	s.Repo.On("FindByID", mock.AnythingOfType("*context.emptyCtx"), want.ID).
		Return(want, nil).Once()

	got, err := s.Service.FindByID(context.Background(), want.ID)

	// Then
	s.Require().NoError(err)
	s.Require().Equal(want, got)
}

func (s *UserServiceTestSuite) TestUserService_FindByUsername() {
	// Given
	want := s.SeedMockUserData()

	// When
	s.Repo.On("FindByUsername", mock.AnythingOfType("*context.emptyCtx"), want.Username).
		Return(want, nil).Once()

	got, err := s.Service.FindByUsername(context.Background(), want.Username)

	// Then
	s.Require().NoError(err)
	s.Require().Equal(want, got)
}

func (s *UserServiceTestSuite) TestUserService_FindByEmail() {
	// Given
	want := s.SeedMockUserData()

	// When
	s.Repo.On("FindByEmail", mock.AnythingOfType("*context.emptyCtx"), want.Email).
		Return(want, nil).Once()

	dbUser, err := s.Service.FindByEmail(context.Background(), want.Email)

	// Then
	s.Require().NoError(err)
	s.Require().Equal(want, dbUser)
}

func (s *UserServiceTestSuite) TestUserService_Delete() {
	// Given
	want := s.SeedMockUserData()

	// When
	s.Repo.On("Delete", mock.AnythingOfType("*context.emptyCtx"), want).
		Return(nil).
		Once()

	err := s.Service.Delete(context.Background(), want)

	// Then
	s.Require().Nil(err)
}
