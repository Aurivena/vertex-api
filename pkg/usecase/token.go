package usecase

import (
	"vertexUP/models"
	"vertexUP/pkg/utils"
)

func (u Usecase) SetToken(login string) (string, utils.ErrorCode) {
	token, err := u.services.GenerateTokenAndSave(login)
	if err != nil {
		return "", utils.BadRequest
	}

	return token.AccessToken, utils.Success
}

func (u Usecase) RefreshAllToken(login string) (*models.Token, utils.ErrorCode) {
	output, err := u.services.RefreshAllToken(login)
	if err != nil {
		return nil, utils.BadRequest
	}
	return output, utils.NoContent
}

func (u Usecase) CheckValidUser(login string) (bool, utils.ErrorCode) {
	_, err := u.services.CheckValidUser(login)
	if err != nil {
		return false, utils.BadRequest
	}

	return true, utils.NoContent
}
