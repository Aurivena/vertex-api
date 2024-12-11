package usecase

import (
	"vertexUP/models"
)

func (u Usecase) SignUp(input *models.SignUpInput) (*models.SignUpOutput, ErrorCode) {
	isRegistered, err := u.services.IsRegistered(input.Login)
	if err != nil {
		return nil, InternalServerError
	}

	if !isRegistered {
		output, err := u.services.SignUp(input)
		if err != nil {
			return nil, BadRequest
		}

		return output, Success
	}
	return nil, BadRequest
}

func (u Usecase) SignIn(input *models.SignInInput) (*models.SignInOutput, ErrorCode) {
	isRegistered, err := u.services.IsRegistered(input.Input)
	if err != nil {
		return nil, InternalServerError
	}

	if isRegistered {
		output, err := u.services.SignIn(input)
		if err != nil {
			return nil, InternalServerError
		}

		return output, Success
	}

	return nil, UnregisteredAccount
}

func (u Usecase) Logout(token string) ErrorCode {
	err := u.services.Logout(token)
	if err != nil {
		return InternalServerError
	}

	return NoContent
}
