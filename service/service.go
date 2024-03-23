package service

import (
	"github.com/Just-Goo/Student-Portal/model"
	"github.com/Just-Goo/Student-Portal/repository"
)

type Service interface {
	CreateStudent(*model.Student)
	GetStudent(*[]model.Student)
	GetStudentById(*model.Student, int) error
	UpdateStudent(*model.Student, *model.Student, int) error
	DeleteStudent(*model.Student, int) error
}

type service struct {
	Service Service
}

func NewService(repo repository.Repository) *service {
	return &service{Service: NewServiceImpl(repo)}
}
