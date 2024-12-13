package usecase

import (
	"vertexUP/models"
)

func (u Usecase) UpdatePassword(input *models.UpdatePasswordInput) error {
	return nil
}

func (u Usecase) GetUserByAccessToken(refreshToken string) (*models.Account, ErrorCode) {
	output, err := u.services.GetUserByAccessToken(refreshToken)
	if err != nil {
		return nil, BadRequest
	}

	return output, Success
}

func (u Usecase) UpdateInfoUser(info *models.UpdateInfoAccountInput, token string) (*models.UpdateInfoAccountOutput, ErrorCode) {

	output, err := u.services.UpdateInfoAccount(info, token)
	if err != nil {
		return nil, BadRequest
	}

	return output, Success
}
