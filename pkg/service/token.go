package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/pkg/errors"
	"time"
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

const (
	access_token_time  = 15 * time.Minute
	refresh_token_time = 14 * 24 * time.Hour
)

type TokenService struct {
	repo   repository.Token
	secret string
}

func NewTokenService(repo repository.Token, secret string) *TokenService {
	return &TokenService{repo: repo, secret: secret}
}

func (s TokenService) GenerateTokenAndSave(login string) (*models.Token, error) {

	accessToken, err := CreateJWTToken(login, s.secret, access_token_time)
	if err != nil {
		return nil, fmt.Errorf("ошибка генерации access token: %w", err)
	}
	refreshToken, err := CreateJWTToken(login, s.secret, refresh_token_time)
	if err != nil {
		return nil, fmt.Errorf("ошибка генерации refresh token: %w", err)
	}
	token := &models.Token{
		AccessToken:           accessToken,
		RefreshToken:          refreshToken,
		AccessTokenExpiresAt:  time.Now().UTC().Add(access_token_time),
		RefreshTokenExpiresAt: time.Now().UTC().Add(refresh_token_time),
	}

	count := s.repo.CheckCount(login)
	if count > 5 {
		s.repo.RevokeToken(accessToken)
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

func (s TokenService) UpdateAccessToken(refreshToken string, login string) (string, error) {
	newAccessToken, err := CreateJWTToken(login, s.secret, access_token_time)
	if err != nil {
		return "", err
	}

	return s.repo.UpdateAccessToken(refreshToken, newAccessToken, time.Now().UTC().Add(access_token_time))
}

func (s TokenService) CheckValidUser(login string) (bool, error) {
	output, err := s.repo.GetTimeToken(login)
	if err != nil {
		return false, err
	}

	if output.AccessTokenExpiresAt.IsZero() {
		return false, errors.New("время истечения access token отсутствует")
	}

	accessTokenDiff := time.Now().UTC().Sub(output.AccessTokenExpiresAt)
	if accessTokenDiff >= access_token_time {
		return false, errors.New("access token истек")
	}

	if output.RefreshTokenExpiresAt.IsZero() {
		return false, errors.New("время истечения refresh token отсутствует")
	}

	refreshTokenDiff := time.Now().UTC().Sub(output.RefreshTokenExpiresAt)
	if refreshTokenDiff > refreshTokenDiff {
		return false, errors.New("refresh token истек")
	}

	return true, nil

}

func (s TokenService) RefreshAllToken(login string) (*models.Token, error) {
	valid, err := s.CheckValidUser(login)
	if err != nil {
		return nil, err
	}

	if !valid {
		return nil, err
	}

	newAccessToken, err := CreateJWTToken(login, s.secret, access_token_time)
	if err != nil {
		return nil, fmt.Errorf("ошибка генерации access token: %w", err)
	}

	newRefreshToken, err := CreateJWTToken(login, s.secret, refresh_token_time)
	if err != nil {
		return nil, fmt.Errorf("ошибка генерации refresh token: %w", err)
	}

	err = s.repo.RefreshAllTokens(login, newAccessToken, newRefreshToken, time.Now().Add(access_token_time), time.Now().Add(refresh_token_time))
	if err != nil {
		return nil, err
	}

	return &models.Token{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func CreateJWTToken(login, secret string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"login": login,
		"exp":   time.Now().UTC().Add(expiration),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
