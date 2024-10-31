package services

import (
	"errors"
	"golang-beginner-19/models"
	"golang-beginner-19/repositories"
)

type MaterialService struct {
	RepoMaterial repositories.MaterialRepositoryDB
}

func NewMaterialService(repo repositories.MaterialRepositoryDB) *MaterialService {
	return &MaterialService{RepoMaterial: repo}
}

func (s *MaterialService) GetMaterialById(id int) (*models.Material, error) {
	if id == 0 {
		return nil, errors.New("id cannot be zero")
	}
	studentFound, err := s.RepoMaterial.GetById(id)
	if err != nil {
		return nil, err
	}
	if studentFound.Status == "deleted" {
		return nil, errors.New("student is deleted")
	}
	return studentFound, nil
}

func (s *MaterialService) GetAllMaterials() ([]models.Material, error) {
	materials, err := s.RepoMaterial.GetAll()
	return *materials, err
}

func (s *MaterialService) CreateMaterial(materialInput *models.Material) error {
	if materialInput == nil {
		return errors.New("material cannot be nil")
	}
	err := s.RepoMaterial.Create(materialInput)
	return err
}
