package services

import (
	"errors"
	"golang-beginner-19/models"
	"golang-beginner-19/repositories"
)

type MentorService struct {
	RepoMentor repositories.MentorRepositoryDB
}

func NewMentorService(repo repositories.MentorRepositoryDB) *MentorService {
	return &MentorService{RepoMentor: repo}
}

func (s *MentorService) UpdateMentor(mentor *models.Mentor) error {
	if mentor.ID == 0 {
		return errors.New("id cannot be zero")
	}
	return s.RepoMentor.Update(mentor)
}

func (s *MentorService) DeleteMentor(id int) error {
	return s.RepoMentor.Delete(id)
}

func (s *MentorService) GetMentorById(id int) (*models.Mentor, error) {
	if id == 0 {
		return nil, errors.New("id cannot be zero")
	}
	mentorFound, err := s.RepoMentor.GetById(id)
	if err != nil {
		return nil, err
	}
	if mentorFound.Status == "deleted" {
		return nil, errors.New("mentor is deleted")
	}
	return mentorFound, nil
}

func (s *MentorService) GetAllMentors() ([]models.Mentor, error) {
	mentors, err := s.RepoMentor.GetAll()
	return *mentors, err
}
