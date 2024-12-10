package usecase

import (
	"vertexUP/models"
	"vertexUP/pkg/utils"
)

func (u Usecase) SignUp(input *models.SignUpInput) (*models.SignUpOutput, utils.ErrorCode) {
	isRegistered, err := u.services.IsRegistered(input.Login)
	if err != nil {
		return nil, utils.InternalServerError
	}

	if !isRegistered {
		output, err := u.services.SignUp(input)
		if err != nil {
			return nil, utils.BadRequest
		}

		return output, utils.Success
	}
	return nil, utils.BadRequest
}

func (u Usecase) SignIn(input *models.SignInInput) (*models.SignInOutput, utils.ErrorCode) {
	isRegistered, err := u.services.IsRegistered(input.Input)
	if err != nil {
		return nil, utils.InternalServerError
	}

	if isRegistered {
		output, err := u.services.SignIn(input)
		if err != nil {
			return nil, utils.InternalServerError
		}

		return output, utils.Success
	}

	return nil, utils.UnregisteredAccount
}

func (u Usecase) Logout(token string) utils.ErrorCode {
	err := u.services.Logout(token)
	if err != nil {
		return utils.InternalServerError
	}

	return utils.NoContent
}
