package handler

import (
	"github.com/gin-gonic/gin"
	"vertexUP/models"
	"vertexUP/pkg/usecase"
)

func (h Handler) updateInfoAccount(c *gin.Context) {

	token := c.GetHeader("Authorization")
	if token == "" {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	input := models.UpdateInfoAccountInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
	}

	output, processStatus := h.usecase.UpdateInfoUser(&input, token)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
	}

	h.sendResponseSuccess(c, output, processStatus)
}
