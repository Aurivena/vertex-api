package service

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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

func (s AccountService) GetUserByAccessToken(accessToken string) (*models.Account, error) {
	return s.repo.GetUserByAccessToken(accessToken)
}

func (s AccountService) UpdateInfoAccount(info *models.UpdateInfoAccountInput, token string) error {
	if info.Email != "" {
		exists := isEmail(info.Email)
		if !exists {
			logrus.Error("введите корректный email!")
			return errors.New("введите корректный email!")
		}
	}
	if info.Password != "" {
		err := validatePassword(info.Password)
		if err != nil {
			return err
		}
		info.Password = bcryptHash(info.Password)
	}

	return s.repo.UpdateInfoUser(info, token)
}
