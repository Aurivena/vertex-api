package models

import "time"

type Account struct {
	Name             string    `json:"name"`
	Login            string    `json:"login"`
	Password         string    `json:"password"`
	Email            string    `json:"email"`
	Status           string    `json:"status"`
	DateRegistration time.Time `json:"date_registration"`
}

type SignInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignInOutput struct {
	Account
}

type SignUpInput struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SignUpOutput struct {
	Name             string    `json:"name"`
	Login            string    `json:"login"`
	Email            string    `json:"email"`
	Status           string    `json:"status"`
	DateRegistration time.Time `json:"date_registration"`
}
