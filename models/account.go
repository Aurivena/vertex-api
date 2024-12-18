package models

import (
	"github.com/google/uuid"
	"time"
)

type Account struct {
	UUID             uuid.UUID `json:"uuid" db:"uuid"`
	Name             string    `json:"name" db:"name"`
	Login            string    `json:"login" db:"login"`
	Password         string    `json:"-"`
	Email            string    `json:"email" db:"email"`
	Status           string    `json:"status" db:"status"`
	DateRegistration time.Time `json:"date_registration" db:"date_registration"`
}

type UpdateInfoAccountInput struct {
	UUID     uuid.UUID `json:"uuid" db:"uuid"`
	Name     string    `json:"name" db:"name"`
	Login    string    `json:"login" db:"login"`
	Password string    `json:"password"`
	Email    string    `json:"email" db:"email"`
}
type UpdateInfoAccountOutput struct {
	Name     string `json:"name" db:"name"`
	Login    string `json:"login" db:"login"`
	Password string `json:"-"`
	Email    string `json:"email" db:"email"`
}

type SignOutput struct {
	Name             string    `json:"name" db:"name"`
	Login            string    `json:"login" db:"login"`
	Email            string    `json:"email" db:"email"`
	Status           string    `json:"status" db:"status"`
	DateRegistration time.Time `json:"date_registration" db:"date_registration"`
}

type SignInInput struct {
	Input    string `json:"input"`
	Password string `json:"password"`
}

type SignInOutput struct {
	SignOutput
}

type SignUpInput struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignUpOutput struct {
	SignOutput
}

type UpdatePasswordInput struct {
	jwt      string `json:"jwt"`
	Password string `json:"password"`
}

type UpdateInfoInput struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
