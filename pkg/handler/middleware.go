package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
	"vertexUP/pkg/usecase"
)

func (h Handler) TokenValidationMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	if token == "" {
		h.sendResponseSuccess(c, nil, usecase.BadHeader)
		return
	}

	outputUser, processStatus := h.usecase.GetUserByAccessToken(token)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	isValid, processStatus := h.usecase.CheckValidUser(outputUser.Login)
	if processStatus != usecase.NoContent {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	if !isValid {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	tokenNew, processStatus := h.usecase.RefreshAllToken(outputUser.Login)
	if processStatus != usecase.NoContent {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}
	c.Request.Header.Set("Authorization", tokenNew.AccessToken)

	c.Next()
}
