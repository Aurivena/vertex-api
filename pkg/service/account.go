package service

import (
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s AccountService) IsRegistered(input string) (bool, error) {
	return s.repo.IsRegistered(input)
}

func (s AccountService) GetUserByEmail(email string) (*models.Account, error) {
	return s.repo.GetUserByEmail(email)
}

func (s AccountService) GetUserByLogin(login string) (*models.Account, error) {
	return s.repo.GetUserByLogin(login)
}
