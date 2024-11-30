package usecase

import (
	"vertexUP/models"
	"vertexUP/pkg/utils"
)

func (u Usecase) SignUp(input *models.SignUpInput) (*models.SignUpOutput, utils.ErrorCode) {
	output, err := u.services.SignUp(input)
	if err != nil {
		return nil, utils.BadRequest
	}

	return output, utils.Success
}
