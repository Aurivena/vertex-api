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
		refreshToken := c.GetHeader("Refresh-Token")

		if refreshToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Требуется авторизация"})
			c.Abort()
			return
		}

		outputUser, processStatus := h.usecase.GetUserByRefreshToken(refreshToken)
		if processStatus != utils.Success {
			c.JSON(http.StatusUnauthorized, gin.H{"error": processStatus})
			return
		}

		tokenNew, processStatus := h.usecase.UpdateAccessToken(refreshToken, outputUser.Login)
		if processStatus != utils.Success {
			c.JSON(http.StatusUnauthorized, gin.H{"error": processStatus})
			return
		}
		c.Request.Header.Set("Authorization", tokenNew)
	} else {
		c.Request.Header.Set("Authorization", token)
	}

	c.Next()
}
