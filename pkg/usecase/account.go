package usecase

import (
	"vertexUP/models"
	"vertexUP/pkg/utils"
)

func (u Usecase) UpdatePassword(input *models.UpdatePasswordInput) error {
	return nil
}

func (u Usecase) GetUserByRefreshToken(refreshToken string) (*models.Account, utils.ErrorCode) {
	output, err := u.services.GetUserByRefreshToken(refreshToken)
	if err != nil {
		return nil, utils.BadRequest
	}

	return output, utils.Success
}
