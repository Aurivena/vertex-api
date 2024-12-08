package usecase

import (
	"vertexUP/pkg/utils"
)

func (u Usecase) SetToken(login string) (string, utils.ErrorCode) {
	token, err := u.services.GenerateTokenAndSave(login)
	if err != nil {
		return "", utils.BadRequest
	}

	return token.AccessToken, utils.Success
}

func (u Usecase) RefreshToken(refreshToken string, login string) (string, utils.ErrorCode) {
	token, err := u.services.RefreshToken(refreshToken, login)
	if err != nil {
		return "", utils.BadRequest
	}

	return token, utils.Success
}
