package usecase

import (
	"vertexUP/models"
)

func (u Usecase) SetToken(login string) (string, ErrorCode) {
	user, err := u.services.GetUserByLogin(login)
	if err != nil {
		return "", BadRequest
	}
	token, err := u.services.GenerateTokenAndSave(user.UUID, user.Login)
	if err != nil {
		return "", BadRequest
	}

	return token.AccessToken, Success
}

func (u Usecase) RefreshAllToken(login string) (*models.Token, ErrorCode) {
	user, err := u.services.GetUserByLogin(login)
	if err != nil {
		return nil, BadRequest
	}
	output, err := u.services.RefreshAllToken(user.UUID, user.Login)
	if err != nil {
		return nil, InternalServerError
	}
	return output, NoContent
}

func (u Usecase) CheckValidUser(login string) (bool, ErrorCode) {
	user, err := u.services.GetUserByLogin(login)
	if err != nil {
		return false, BadRequest
	}
	_, err = u.services.CheckValidUser(user.UUID)
	if err != nil {
		return false, InternalServerError
	}

	return true, NoContent
}
