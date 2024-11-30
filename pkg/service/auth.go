package service

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"vertexUP/clerr"
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

type AuthService struct {
	repo repository.Auth
}

func NewAuthService(repo repository.Auth) *AuthService { return &AuthService{repo: repo} }

func (s AuthService) SignIn(input *models.SignInInput) (*models.SignInOutput, error) {
	return nil, nil
}

func (s AuthService) SignUp(input *models.SignUpInput) (*models.SignUpOutput, error) {
	if err := validateName(input.Name); err != nil {
		return nil, err
	}

	if err := validateLogin(input.Login); err != nil {
		return nil, err
	}

	if err := validateEmail(input.Email); err != nil {
		return nil, err
	}

	if err := validatePassword(input.Password); err != nil {
		return nil, err
	}

	return s.repo.SignUp(input)
}

func (s AuthService) SignOut() {

}

func validateName(name string) error {
	if name == "" {
		logrus.Error("name is nil")
		return errors.New("name is nil")
	}
	return nil
}

func validateLogin(login string) error {
	if login == "" {
		logrus.Error("login is nil")
		return errors.New("login is nil")
	}
	return nil
}

func validateEmail(email string) error {
	if email == "" {
		logrus.Error("email is nil")
		return errors.New("email is nil")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		logrus.Error("password must be at least 8 characters")
		return errors.New(clerr.ErrorPasswordTooShort)
	}
	return nil
}
