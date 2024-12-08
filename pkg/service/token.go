package service

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

type TokenService struct {
	repo   repository.Repository
	secret string
}

func NewTokenService(repo repository.Repository, secret string) *TokenService {
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

	err = s.repo.SaveToken(login, *token)
	if err != nil {
		return nil, err
	}

	return token, nil
}
