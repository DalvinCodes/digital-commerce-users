package config

import (
	"fmt"
	"github.com/DalvinCodes/digital-commerce/users/model"
	zap2 "go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"moul.io/zapgorm2"
)

type UsersDB struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func NewUsersDatabase(configFile *Configurations) *gorm.DB {
	url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		configFile.Postgres.Username,
		configFile.Postgres.Password,
		configFile.Postgres.Host,
		configFile.Postgres.Port,
		configFile.Postgres.Name)

	zap := zapgorm2.New(zap2.L())
	zap.SetAsDefault()

	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger:                 zap,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	if err := pingDB(db); err != nil {
		log.Fatal(err)
	}

	return migrateDB(db)
}

func migrateDB(db *gorm.DB) *gorm.DB {
	if err := db.AutoMigrate(
		//TODO: add new models to migrator

		&model.User{},
		&model.Address{},
	); err != nil {
		log.Fatalln(err)
	}
	return db
}

func pingDB(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	log.Print("Ping Successful")
	return sqlDB.Ping()
}
