package repository

import (
	"fmt"

	"github.com/Just-Goo/Student-Portal/model"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	DB *gorm.DB
}

func (r *RepositoryImpl) Create(student *model.Student)  {
	r.DB.Create(student)
}

func (r *RepositoryImpl) GetById(student *model.Student, id int) error  {
	r.DB.Find(student, "id = ?", id)
	if student.ID == 0 {
		return fmt.Errorf("student does not exist")
	}
	return nil
}

func (r *RepositoryImpl) Get(students *[]model.Student) {
	r.DB.Find(students)
}

func (r *RepositoryImpl) Update(student *model.Student)  {
	r.DB.Save(student)
}

func (r *RepositoryImpl) Delete(student *model.Student) error {
	return r.DB.Delete(student).Error
}

func NewRepositoryImpl(db *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{DB: db}
}