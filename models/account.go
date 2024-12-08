package models

import "time"

type Account struct {
	Name             string    `json:"name"`
	Login            string    `json:"login"`
	Token            Token     `json:"token"`
	Password         string    `json:"-"`
	Email            string    `json:"email"`
	Status           string    `json:"status"`
	DateRegistration time.Time `json:"date_registration" db:"date_registration"`
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
