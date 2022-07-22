package main

import (
	"github.com/DalvinCodes/digital-commerce/users/config"
	"github.com/DalvinCodes/digital-commerce/users/repo"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	config.Vars = config.LoadConfigs()

	database := config.NewUsersDatabase(config.Vars)
	repo.NewUserRepository(database)
	server := fiber.New()
	log.Fatal(server.Listen(config.Vars.Server.Port))
}
