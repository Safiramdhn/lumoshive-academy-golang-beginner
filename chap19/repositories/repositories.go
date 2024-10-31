package repositories

import "golang-beginner-19/models"

type UserRepository interface {
	Create(string) (int, error)
	Login(email, password string) (bool, error)
}

type SchduleRepository interface {
	Create(schedule *models.Schedule) error
	GetAll() (*[]models.Schedule, error)
}

type MaterialRepository interface {
	Create(material *models.Material) error
	GetAll() (*[]models.Material, error)
}

type StudentRepository interface {
	GetAll() (*[]models.Student, error)
	GetById(id int) (*models.Student, error)
	Update(student *models.Student) error
	Delete(id int) error
}

// type MentorRepository interface {
// }
