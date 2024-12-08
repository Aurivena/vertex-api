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
		c.JSON(http.StatusUnauthorized, gin.H{"error": "токен отсутствует"})
		c.Abort()
		return
	}

	isActive, err := h.usecase.IsTokenActive(token)
	if err != utils.Success {
		refreshToken := c.GetHeader("Refresh-Token")
		if refreshToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Требуется авторизация"})
			return
		}

		outputUser, processStatus := h.usecase.GetUserByRefreshToken(refreshToken)
		if processStatus != utils.Success {
			c.JSON(http.StatusUnauthorized, gin.H{"error": processStatus})
			return
		}

		tokenNew, processStatus := h.usecase.RefreshToken(refreshToken, outputUser.Login)
		if processStatus != utils.Success {
			c.JSON(http.StatusUnauthorized, gin.H{"error": processStatus})
			return
		}

		c.Request.Header.Set("Authorization", tokenNew)
	} else {
		c.Request.Header.Set("Authorization", token)
	}

	if !isActive {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "токен недействителен"})
		c.Abort()
		return
	}

	c.Next()
}
