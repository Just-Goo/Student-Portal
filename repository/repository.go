package repository

import (
	"github.com/Just-Goo/Student-Portal/model"
	"gorm.io/gorm"
)

type Repository interface {
	Create(*model.Student)
	Get(*[]model.Student)
	GetById(*model.Student, int) error
	Update(*model.Student)
	Delete(*model.Student) error
}

type repository struct {
	Repo Repository
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{Repo: NewRepositoryImpl(db)}
}
