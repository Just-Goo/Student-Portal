package service

import (
	"github.com/Just-Goo/Student-Portal/model"
	"github.com/Just-Goo/Student-Portal/repository"
)

type ServiceImpl struct {
	Repo repository.Repository
}

// GetStudent(*[]model.Student)
func (s *ServiceImpl) CreateStudent(student *model.Student) {
	s.Repo.Create(student)
}

func (s *ServiceImpl) GetStudent(students *[]model.Student) {
	s.Repo.Get(students)
}

func (s *ServiceImpl) DeleteStudent(student *model.Student, id int) error {
	if err := s.Repo.GetById(student, id); err != nil {
		return err
	}

	return s.Repo.Delete(student)
}

func (s *ServiceImpl) UpdateStudent(student, updatedStudent *model.Student, id int) error {
	// Get student from database
	if err := s.Repo.GetById(student, id); err != nil {
		return err
	}

	// update the required information
	student.FirstName = updatedStudent.FirstName
	student.LastName = updatedStudent.LastName
	student.Age = updatedStudent.Age
	student.Gender = updatedStudent.Gender

	s.Repo.Update(student)
	return nil
}

func (s *ServiceImpl) GetStudentById(student *model.Student, id int) error {
	return s.Repo.GetById(student, id)
}

func NewServiceImpl(repo repository.Repository) *ServiceImpl {
	return &ServiceImpl{Repo: repo}
}
