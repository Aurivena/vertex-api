package handler

import (
	"github.com/gin-gonic/gin"
	"vertexUP/models"
	"vertexUP/pkg/usecase"
)

// @Summary      Обновить данные пользователя
// @Description  Обновляет данные пользователя, которые будут отправлены
// @Tags         Аккаунт
// @Accept       json
// @Produce      json
// @Param        models.UpdateInfoAccountInput body models.UpdateInfoAccountInput true "Входные данные"
// @Success      204 {object} string "NoContent"
// @Failure      400 {object} string "BadRequest"
// @Router       /account [put]
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

	processStatus := h.usecase.UpdateInfoUser(&input, token)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
	}

	h.sendResponseSuccess(c, nil, processStatus)
}
