package utils

import "github.com/google/uuid"

func GenerateUUID() (uuid.UUID, error) {
	output, err := uuid.NewV7()
	if err != nil {
		return uuid.Nil, err
	}
	return output, nil
}
