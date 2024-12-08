package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

type TokenService struct {
	repo   repository.Token
	secret string
}

func NewTokenService(repo repository.Token, secret string) *TokenService {
	return &TokenService{repo: repo, secret: secret}
}

func (s TokenService) GenerateTokenAndSave(login string) (*models.Token, error) {

	now := time.Now()

	accessTokenClaims := jwt.MapClaims{
		"login": login,
		"exp":   now.Add(15 * time.Minute).Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	refreshTokenClaims := jwt.MapClaims{
		"login": login,
		"exp":   now.Add(7 * 24 * time.Hour).Unix(),
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.secret))
	if err != nil {
		return nil, err
	}

	token := &models.Token{
		AccessToken:           accessTokenString,
		RefreshToken:          refreshTokenString,
		AccessTokenExpiresAt:  now.Add(15 * time.Minute),
		RefreshTokenExpiresAt: now.Add(7 * 24 * time.Hour),
	}

	count := s.repo.CheckCount(login)
	if count > 5 {
		s.repo.RevokeToken(accessTokenString)
		return nil, fmt.Errorf("ошибка. токенов больше 5")

	}

	err = s.repo.SaveToken(login, *token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s TokenService) Logout(token string) error {
	err := s.repo.RevokeToken(token)
	if err != nil {
		return err
	}
	return nil
}
