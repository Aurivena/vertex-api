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

		token, err := u.services.GenerateTokenAndSave(output.Login)
		if err != nil {
			return nil, utils.InternalServerError
		}
		output.Token = models.Token{
			AccessToken:           token.AccessToken,
			RefreshToken:          token.RefreshToken,
			AccessTokenExpiresAt:  token.AccessTokenExpiresAt,
			RefreshTokenExpiresAt: token.RefreshTokenExpiresAt,
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

		token, err := u.services.GenerateTokenAndSave(output.Login)
		if err != nil {
			return nil, utils.InternalServerError
		}
		output.Token = models.Token{
			AccessToken:           token.AccessToken,
			RefreshToken:          token.RefreshToken,
			AccessTokenExpiresAt:  token.AccessTokenExpiresAt,
			RefreshTokenExpiresAt: token.RefreshTokenExpiresAt,
		}

		return output, utils.Success
	}

	return nil, utils.UnregisteredAccount
}
