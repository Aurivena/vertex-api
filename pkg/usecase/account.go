package usecase

import (
	"vertexUP/models"
	"vertexUP/pkg/utils"
)

func (u Usecase) UpdatePassword(input *models.UpdatePasswordInput) error {
	return nil
}

func (u Usecase) GetUserByAccessToken(refreshToken string) (*models.Account, utils.ErrorCode) {
	output, err := u.services.GetUserByAccessToken(refreshToken)
	if err != nil {
		return nil, utils.BadRequest
	}

	return output, utils.Success
}
