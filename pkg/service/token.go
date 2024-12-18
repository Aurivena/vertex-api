package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"time"
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

const (
	access_token_time     = 15 * time.Minute
	refresh_token_time    = 14 * 24 * time.Hour
	max_inactive_duration = 6 * 24 * time.Hour
)

type TokenService struct {
	repo   repository.Token
	secret string
}

func NewTokenService(repo repository.Token, secret string) *TokenService {
	return &TokenService{repo: repo, secret: secret}
}

func (s TokenService) GenerateTokenAndSave(uuid uuid.UUID, login string) (*models.Token, error) {

	err := s.repo.CheckCount(uuid)
	if err != nil {
		logrus.Errorf("Ошибка при проверке количества токенов для пользователя %s: %v", login, err)
		return nil, err
	}

	accessToken, err := CreateJWTToken(login, s.secret, access_token_time)
	if err != nil {
		logrus.Errorf("Ошибка при генерации access токена для пользователя %s: %v", login, err)
		return nil, fmt.Errorf("ошибка генерации access token: %w", err)
	}
	refreshToken, err := CreateJWTToken(login, s.secret, refresh_token_time)
	if err != nil {
		logrus.Errorf("Ошибка при генерации refresh токена для пользователя %s: %v", login, err)
		return nil, fmt.Errorf("ошибка генерации refresh token: %w", err)
	}
	token := &models.Token{
		AccessToken:         accessToken,
		RefreshToken:        refreshToken,
		AccessTokenExpires:  time.Now().UTC().Add(access_token_time),
		RefreshTokenExpires: time.Now().UTC().Add(refresh_token_time),
	}

	err = s.repo.SaveToken(uuid, *token)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s TokenService) Logout(token string) error {
	err := s.repo.DeleteToken(token)
	if err != nil {
		return err
	}
	return nil
}

func (s TokenService) CheckValidUser(uuid uuid.UUID) (bool, error) {
	output, err := s.repo.GetAllInfoToken(uuid)
	if err != nil {
		return false, err
	}

	accessTokenDiff := time.Now().UTC().Sub(output.AccessTokenExpires)

	if accessTokenDiff > max_inactive_duration {
		return false, nil
	}

	return true, nil

}

func (s TokenService) RefreshAllToken(uuid uuid.UUID, login string) (*models.Token, error) {
	currentTokens, err := s.repo.GetAllInfoToken(uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить текущие токены: %w", err)
	}

	remainingAccessTime := time.Until(currentTokens.AccessTokenExpires)
	remainingRefreshTime := time.Until(currentTokens.RefreshTokenExpires)

	if remainingRefreshTime < 0 && remainingAccessTime < 0 {
		return nil, errors.New("ваша сессия истекла. авторизуйтесь заново")
	}

	var newAccessToken string
	if remainingAccessTime < 0 {
		newAccessToken, err = CreateJWTToken(login, s.secret, access_token_time)
		if err != nil {
			return nil, fmt.Errorf("ошибка генерации access_token: %w", err)
		}
		err = s.repo.UpdateAccessToken(currentTokens.AccessToken, newAccessToken, time.Now().Add(access_token_time))
		if err != nil {
			return nil, fmt.Errorf("ошибка обновления access_token: %w", err)
		}
	} else {
		newAccessToken = currentTokens.AccessToken
	}

	var newRefreshToken string
	if remainingRefreshTime < refresh_token_time/2 {
		newRefreshToken, err = CreateJWTToken(login, s.secret, refresh_token_time)
		if err != nil {
			return nil, fmt.Errorf("ошибка генерации refresh_token: %w", err)
		}
		err = s.repo.UpdateRefreshToken(currentTokens.RefreshToken, newRefreshToken, time.Now().Add(refresh_token_time))
		if err != nil {
			return nil, fmt.Errorf("ошибка обновления refresh_token: %w", err)
		}
	} else {
		newRefreshToken = currentTokens.RefreshToken
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

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return signedToken + "-" + uuid.New().String(), nil
}
