package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"vertexUP/pkg/utils"
)

func (h Handler) TokenValidationMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	token = strings.TrimPrefix(token, "Bearer ")

	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется авторизация"})
		c.Abort()
		return
	}

	outputUser, processStatus := h.usecase.GetUserByAccessToken(token)
	if processStatus != utils.Success {
		c.JSON(http.StatusUnauthorized, gin.H{"error": processStatus})
		c.Abort()
		return
	}

	isValid, processStatus := h.usecase.CheckValidUser(outputUser.Login)
	if processStatus != utils.NoContent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": processStatus})
		c.Abort()
		return
	}

	if !isValid {
		c.JSON(http.StatusInternalServerError, gin.H{"error": processStatus})
		c.Abort()
		return
	}

	tokenNew, processStatus := h.usecase.RefreshAllToken(outputUser.Login)
	if processStatus != utils.NoContent {
		c.JSON(http.StatusInternalServerError, gin.H{"error": processStatus})
		c.Abort()
		return
	}
	c.Request.Header.Set("Authorization", tokenNew.AccessToken)

	c.Next()
}
