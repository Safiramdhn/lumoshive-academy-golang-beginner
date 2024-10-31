package services

import (
	"errors"
	"golang-beginner-19/models"
	"golang-beginner-19/repositories"
)

type StudentService struct {
	RepoStudent repositories.StudentRepositoryDB
}

func NewStudentService(repo repositories.StudentRepositoryDB) *StudentService {
	return &StudentService{RepoStudent: repo}
}

func (s *StudentService) UpdateStudent(student *models.Student) error {
	if student.ID == 0 {
		return errors.New("id cannot be zero")
	}
	return s.RepoStudent.Update(student)
}

func (s *StudentService) DeleteStudent(id int) error {
	return s.RepoStudent.Delete(id)
}

func (s *StudentService) GetStudentById(id int) (*models.Student, error) {
	if id == 0 {
		return nil, errors.New("id cannot be zero")
	}
	studentFound, err := s.RepoStudent.GetById(id)
	if err != nil {
		return nil, err
	}
	if studentFound.Status == "deleted" {
		return nil, errors.New("student is deleted")
	}
	return studentFound, nil
}

func (s *StudentService) GetAllStudents() ([]models.Student, error) {
	students, err := s.RepoStudent.GetAll()
	return *students, err
}
