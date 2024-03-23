package controller

import (
	"github.com/Just-Goo/Student-Portal/service"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateStudent(c *fiber.Ctx) error
	GetStudents(c *fiber.Ctx) error
	GetSingleStudent(c *fiber.Ctx) error
	DeleteStudent(c *fiber.Ctx) error
	UpdateStudent(c *fiber.Ctx) error
}

type handler struct {
	Handler Handler
}

func NewHandler(service service.Service) *handler {
	return &handler{Handler: NewHandlerImpl(service)}
}
