package service

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

type AuthService struct {
	repo    repository.Auth
	account *AccountService
}

func NewAuthService(repo repository.Auth, account *AccountService) *AuthService {
	return &AuthService{repo: repo, account: account}
}

func (s AuthService) SignIn(input *models.SignInInput) (*models.SignInOutput, error) {
	var user *models.Account
	var err error
	if isEmail(input.Input) {
		user, err = s.account.GetUserByEmail(input.Input)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = s.account.GetUserByLogin(input.Input)
		if err != nil {
			return nil, err
		}
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, err
	}

	return s.repo.SignIn(input)
}

func (s AuthService) SignUp(input *models.SignUpInput) (*models.SignUpOutput, error) {
	if err := validateName(input.Name); err != nil {
		return nil, err
	}

	if err := validateLogin(input.Login); err != nil {
		return nil, err
	}

	if !isEmail(input.Email) {
		return nil, errors.Errorf("email %s не валиден", input.Email)
	}

	if err := validatePassword(input.Password); err != nil {
		return nil, err
	}

	input.Password = bcryptHash(input.Password)

	return s.repo.SignUp(input, time.Now().UTC())
}

func validateName(name string) error {
	if name == "" {
		logrus.Error("имя пустое")
		return errors.New("имя пустое")
	}
	return nil
}

func validateLogin(login string) error {
	if login == "" {
		logrus.Error("логин пустой")
		return errors.New("логин пустой")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		logrus.Error("у пароля длина меньше 8 символов")
		return errors.New("у пароля длина меньше 8 символов")
	}
	return nil
}

func isEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func bcryptHash(password string) string {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	return string(passwordHash)
}
