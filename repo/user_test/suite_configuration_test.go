package repo

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/DalvinCodes/digital-commerce/users/model"
	"github.com/DalvinCodes/digital-commerce/users/repo"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

type UserTestSuite struct {
	suite.Suite
	Repo     repo.UserRepo
	Mock     sqlmock.Sqlmock
	DB       *sql.DB
	User     model.User
	UserList []*model.User
}

func (s *UserTestSuite) SetupSuite() {
	s.T().Log("Setting Up User Test Suite.")

	//Setting up Mock DB and Mock Test Expectation Suite
	var db *sql.DB
	var err error

	db, s.Mock, err = sqlmock.New()
	if err != nil {
		s.T().Logf("error setting up mock database suite: %v", err)
		s.FailNow(err.Error())
	}

	//selecting postgres as base DB provider -- dummy DSN
	dialector := postgres.New(postgres.Config{
		DSN:        "sqlmock_db_0",
		DriverName: "postgres",
		Conn:       db})

	gormDB, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		s.T().Log("error establishing gormORM db connection")
		s.FailNow(err.Error())
	}

	s.Repo.Db = gormDB

	s.DB, err = gormDB.DB()
	if err != nil {
		s.T().Log("error returning base DB interface from gormORM")
		s.FailNow(err.Error())
	}
	s.T().Log("User Test Suite setup is successful.")
}

func (s *UserTestSuite) SeedUser() *model.User {
	user := s.User
	user.ID = gofakeit.UUID()
	user.Username = gofakeit.Username()
	user.FirstName = gofakeit.FirstName()
	user.LastName = gofakeit.LastName()
	user.Email = gofakeit.Email()
	user.DateOfBirth = gofakeit.Date().Format("01/02/2006")

	return &user
}

func (s *UserTestSuite) SeedUserList() {
	for i := 0; i < 10; i++ {
		user := s.SeedUser()
		s.UserList = append(s.UserList, user)
	}
}

func TestRunUserTestSuite(t *testing.T) {
	suite.Run(t, &UserTestSuite{})
}
