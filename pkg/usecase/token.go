package usecase

import (
	"vertexUP/models"
)

func (u Usecase) SetToken(login string) (string, ErrorCode) {
	token, err := u.services.GenerateTokenAndSave(login)
	if err != nil {
		return "", BadRequest
	}

	return token.AccessToken, Success
}

func (u Usecase) RefreshAllToken(login string) (*models.Token, ErrorCode) {
	output, err := u.services.RefreshAllToken(login)
	if err != nil {
		return nil, BadRequest
	}
	return output, NoContent
}

func (u Usecase) CheckValidUser(login string) (bool, ErrorCode) {
	_, err := u.services.CheckValidUser(login)
	if err != nil {
		return false, BadRequest
	}

	return true, NoContent
}
