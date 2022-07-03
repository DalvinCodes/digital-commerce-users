package main

import (
	"github.com/DalvinCodes/digital-commerce/users/config"
	"github.com/DalvinCodes/digital-commerce/users/repo"
	"github.com/gofiber/fiber/v2"
	"log"
)

var Configurations = config.LoadConfigs()

func main() {
	database := config.NewUsersDatabase(Configurations)
	repo.NewUserRepository(database)
	server := fiber.New()
	log.Fatal(server.Listen(Configurations.Server.Port))
}
