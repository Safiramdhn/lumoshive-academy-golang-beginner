package services

import (
	"errors"
	"golang-beginner-19/models"
	"golang-beginner-19/repositories"
)

type UserService struct {
	RepoUser repositories.UserRepositoryDB
}

func NewUserService(repoUser repositories.UserRepositoryDB) *UserService {
	return &UserService{RepoUser: repoUser}
}

func (s *UserService) CreateUser(email, password, first_name, last_name, role string, added_by int) error {
	if email == "" || password == "" || first_name == "" || last_name == "" || role == "" {
		return errors.New("all fields are required")
	}

	err := s.RepoUser.Create(email, password, first_name, last_name, role, added_by)
	if err != nil {
		return err
	}
	return nil
}

func (s *UserService) UpdateUser(user *models.User) error {
	if user.ID == 0 {
		return errors.New("user id is required")
	}
	return s.RepoUser.Update(user)
}
