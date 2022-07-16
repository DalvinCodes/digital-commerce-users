package repo

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"gorm.io/gorm"
	"regexp"
)

func (s *UserTestSuite) TestUserRepo_NewRepository() {
	// Given
	var gormDB *gorm.DB

	// When
	got := NewUserRepository(gormDB)
	want := NewUserRepository(gormDB)

	// Then
	s.Require().Equalf(got, want, "Got %v : Want: %v ", got, want)
}

func (s *UserTestSuite) TestUserRepo_Create() {
	// Given
	user := s.SeedMockUserData()
	const userQuery = `INSERT INTO "users" ("id","username","first_name","last_name","email","dob","created_at","updated_at","deleted_at") VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9) RETURNING "id"`

	rows := s.Mock.NewRows([]string{"id"}).AddRow(user.ID)

	// When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WithArgs(
			user.ID, user.Username, user.FirstName,
			user.LastName, user.Email, user.DateOfBirth,
			sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

	if err := s.Repo.Create(context.Background(), user); err != nil {
		s.Require().Nil(err)
	}

	// Then
	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUserRepo_ListAll() {
	// Given
	const userQuery = `SELECT * FROM "users" WHERE "users"."deleted_at" IS NULL`

	users := s.createUserList()

	rows := s.Mock.NewRows(
		[]string{
			"id", "username", "first_name", "last_name", "email", "dob"}).
		AddRow(users[0].ID, users[0].Username, users[0].FirstName, users[0].LastName, users[0].Email, users[0].DateOfBirth).
		AddRow(users[1].ID, users[1].Username, users[1].FirstName, users[1].LastName, users[1].Email, users[1].DateOfBirth).
		AddRow(users[2].ID, users[2].Username, users[2].FirstName, users[2].LastName, users[2].Email, users[2].DateOfBirth).
		AddRow(users[3].ID, users[3].Username, users[3].FirstName, users[3].LastName, users[3].Email, users[3].DateOfBirth).
		AddRow(users[4].ID, users[4].Username, users[4].FirstName, users[4].LastName, users[4].Email, users[4].DateOfBirth)

	// When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WillReturnRows(rows)

	actualUsers, err := s.Repo.ListAll(context.Background())

	// Then
	s.Require().NoError(err, "error calling db for ListAll: %v", err)
	s.Require().NotNil(actualUsers)
	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUserRepo_ListAll_ReturnsError() {
	// Given
	const userQuery = `SELECT * FROM "users"`

	// When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WillReturnError(errors.New("unable to return a collection of users"))
	user, err := s.Repo.ListAll(context.Background())

	// Then
	s.Require().Error(err, "error was expected while retrieving all users")
	s.Require().Nil(user)

	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUserRepo_FindByID() {
	// Given
	const userQuery = `SELECT * FROM "users" WHERE id = $1`
	mockUser := s.SeedMockUserData()

	rows := s.Mock.NewRows([]string{"id", "username", "first_name", "last_name", "email", "dob"}).
		AddRow(mockUser.ID, mockUser.Username, mockUser.FirstName, mockUser.LastName, mockUser.Email, mockUser.DateOfBirth)

	// When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WithArgs(mockUser.ID).
		WillReturnRows(rows)

	// Then
	dbUser, err := s.Repo.FindByID(context.Background(), mockUser.ID)
	s.Require().NoError(err, "unexpected error while creating dbUser")
	s.Require().Equal(mockUser, dbUser)
	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUserRepo_FindByID_ReturnsError() {
	// Given
	const userQuery = `SELECT * FROM "users" WHERE id = $1`
	obj := s.SeedMockUserData()

	// When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WithArgs(obj.ID).
		WillReturnError(errors.New("unable to query db for user"))

	user, err := s.Repo.FindByID(context.Background(), obj.ID)

	// Then
	s.Require().Error(err, "error was expected while retrieving user")
	s.Require().Empty(user)

	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUserRepo_Update() {
	// Given
	const userQuery = `UPDATE "users" SET "username"=$1,"first_name"=$2,"last_name"=$3,"email"=$4,"dob"=$5,"updated_at"=$6 WHERE "users"."deleted_at" IS NULL AND "id" = $7`
	user := s.SeedMockUserData()

	// When
	s.Mock.ExpectExec(regexp.QuoteMeta(userQuery)).
		WithArgs(
			user.Username, user.FirstName, user.LastName,
			user.Email, user.DateOfBirth, sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := s.Repo.Update(context.Background(), user)

	//  Then
	s.Require().Nil(err)
	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUserRepo_Delete() {
	// Given
	user := s.SeedMockUserData()
	deleteUserQuery := `UPDATE "users" SET "deleted_at"=$1 WHERE "users"."id" = $2 AND "users"."deleted_at" IS NULL`

	// When
	s.Mock.ExpectExec(regexp.QuoteMeta(deleteUserQuery)).
		WithArgs(sqlmock.AnyArg(), user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := s.Repo.Delete(context.Background(), user)
	s.Require().NoError(err)

	s.T().Log(user)

	// Then
	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUserRepo_FindByIDReturnsNoUserFoundError() {
	// Given
	const findUserQuery = `SELECT * FROM "users" WHERE id = $1`
	queryUser := s.SeedMockUserData()

	// When
	s.Mock.ExpectQuery(regexp.QuoteMeta(findUserQuery)).WithArgs(queryUser.ID).
		WillReturnError(errors.New("no user found"))

	user, err := s.Repo.FindByID(context.Background(), queryUser.ID)

	// Then
	s.Require().Error(err)
	s.Require().Empty(user)

	errExpectations := s.Mock.ExpectationsWereMet()
	s.Require().Nil(errExpectations)
}

func (s *UserTestSuite) TestUserRepo_FindByUsername() {
	// Given
	const userQuery = `SELECT * FROM "users" WHERE username = $1 AND "users"."deleted_at" IS NULL`
	user := s.SeedMockUserData()

	rows := s.Mock.NewRows([]string{"id", "username", "first_name", "last_name", "email", "dob"}).
		AddRow(user.ID, user.Username, user.FirstName, user.LastName, user.Email, user.DateOfBirth)

	// When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WithArgs(user.Username).
		WillReturnRows(rows)

	dbUser, err := s.Repo.FindByUsername(context.Background(), user.Username)
	s.Require().NoError(err)

	// Then
	s.Require().Equal(user, dbUser)
	err = s.Mock.ExpectationsWereMet()
	s.Require().NoError(err)
}

func (s *UserTestSuite) TestUserRepo_FindByUsername_ReturnsError() {
	// Given
	const userQuery = `SELECT * FROM "users" WHERE username = $1 AND "users"."deleted_at" IS NULL`
	user := s.SeedMockUserData()

	// When
	s.Mock.ExpectQuery(regexp.QuoteMeta(userQuery)).
		WithArgs(user.Username).
		WillReturnError(errors.New("unable to return user by username"))

	dbUser, err := s.Repo.FindByUsername(context.Background(), user.Username)

	// Then
	s.Require().Error(err)
	s.Require().Empty(dbUser)
	err = s.Mock.ExpectationsWereMet()
	s.Require().NoError(err)
}

func (s *UserTestSuite) createUserList() (users []*model.User) {
	for i := 0; i < 5; i++ {
		users = append(users, s.SeedMockUserData())
	}
	return
}
