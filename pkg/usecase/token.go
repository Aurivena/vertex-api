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
