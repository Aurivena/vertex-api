package service

import "vertexUP/pkg/repository"

type AuthService struct {
	repo repository.Auth
}

func (a AuthService) SignIn() {

}

func (a AuthService) SignUp() {

}

func (a AuthService) SignOut() {

}

func NewAuthService(repo repository.Auth) *AuthService { return &AuthService{repo: repo} }
