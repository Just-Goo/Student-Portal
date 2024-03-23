package main

import (
	"log"

	"github.com/Just-Goo/Student-Portal/controller"
	"github.com/Just-Goo/Student-Portal/database"
	"github.com/Just-Goo/Student-Portal/repository"
	"github.com/Just-Goo/Student-Portal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()

	// dependency injection
	repo := repository.NewRepository(db)
	service := service.NewService(repo.Repo)

	controller.RegisterRoutes(app, service.Service)

	log.Fatal(app.Listen(":3000"))

}
