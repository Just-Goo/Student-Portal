package controller

import (
	"github.com/Just-Goo/Student-Portal/service"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, service service.Service) {

	handler := NewHandler(service)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome")
	})

	v1 := app.Group("/api/v1")

	v1.Post("/student", handler.Handler.CreateStudent)
	v1.Get("/students", handler.Handler.GetStudents)
	v1.Get("/students/:id", handler.Handler.GetSingleStudent)
	v1.Put("/students/:id", handler.Handler.UpdateStudent)
	v1.Delete("/students/:id", handler.Handler.DeleteStudent)

}
