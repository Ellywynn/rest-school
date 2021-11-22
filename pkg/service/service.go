package service

import "github.com/ellywynn/rest-school/repository"

type Service struct {
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
