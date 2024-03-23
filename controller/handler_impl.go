package controller

import (
	"github.com/Just-Goo/Student-Portal/model"
	"github.com/Just-Goo/Student-Portal/service"
	"github.com/gofiber/fiber/v2"
)

type HandlerImpl struct {
	Service service.Service
}

func (h *HandlerImpl) CreateStudent(c *fiber.Ctx) error {
	var student model.Student
	if err := c.BodyParser(&student); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	h.Service.CreateStudent(&student)
	return c.Status(fiber.StatusOK).JSON(student)

}

func (h *HandlerImpl) GetStudents(c *fiber.Ctx) error {
	var students []model.Student
	h.Service.GetStudent(&students)

	return c.Status(fiber.StatusOK).JSON(students)
}

func (h *HandlerImpl) GetSingleStudent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be an integer")
	}

	var student model.Student

	if err = h.Service.GetStudentById(&student, id); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).JSON(student)
}

func (h *HandlerImpl) UpdateStudent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be an integer")
	}

	var student, updatedStudent model.Student

	if err := c.BodyParser(&updatedStudent); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	h.Service.UpdateStudent(&student, &updatedStudent, id)
	return c.Status(fiber.StatusOK).JSON(student)
}

func (h *HandlerImpl) DeleteStudent(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("id must be an integer")
	}
	
	var student model.Student
	if err = h.Service.DeleteStudent(&student, id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return c.Status(fiber.StatusOK).SendString("student deleted successfully")
}

func NewHandlerImpl(service service.Service) *HandlerImpl {
	return &HandlerImpl{Service: service}
}
