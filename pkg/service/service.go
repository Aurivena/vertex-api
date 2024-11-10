package service

import (
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

type Service struct {
}

func NewService(repo *repository.Repository, cfg *models.Config, env *models.Environment) *Service {
	return &Service{}
}
