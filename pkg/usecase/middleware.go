package usecase

import (
	"strings"
	"vertexUP/pkg/utils"
)

func (u Usecase) IsTokenActive(token string) (bool, utils.ErrorCode) {
	token = strings.TrimPrefix(token, "Bearer ")
	isActive, err := u.services.IsTokenActive(token)
	if err != nil {
		return false, utils.BadHeader
	}

	return isActive, utils.Success
}
