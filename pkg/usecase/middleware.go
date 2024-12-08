package usecase

import (
	"vertexUP/pkg/utils"
)

func (u Usecase) IsTokenActive(token string) (bool, utils.ErrorCode) {
	isActive, err := u.services.IsTokenActive(token)
	if err != nil {
		return false, utils.BadHeader
	}

	return isActive, utils.Success
}
