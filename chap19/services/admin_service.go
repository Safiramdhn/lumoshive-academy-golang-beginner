package services

import (
	"errors"
	"fmt"
	"golang-beginner-19/repositories"
)

type AdminService struct {
	RepoAdmin repositories.AdminRepositoryDB
}

func NewAdminService(repoAdmin repositories.AdminRepositoryDB) *AdminService {
	return &AdminService{RepoAdmin: repoAdmin}
}

func (s *AdminService) LoginService(email, password string) (int, error) {
	if email == "" || password == "" {
		return 0, errors.New("email and password are required")
	}

	fmt.Printf("Login request: email=%s, password=%s\n", email, password)
	adminId, err := s.RepoAdmin.Login(email, password)
	if err != nil {
		return 0, err
	}

	if adminId == 0 {
		return 0, errors.New("invalid email or password")
	}
	return adminId, nil
}
