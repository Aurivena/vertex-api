package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vertexUP/pkg/utils"
)

func (h Handler) TokenValidationMiddleware(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "токен отсутствует"})
		c.Abort()
		return
	}

	isActive, err := h.usecase.IsTokenActive(token)
	if err != utils.Success {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ошибка проверки токена"})
		c.Abort()
		return
	}

	c.Request.Header.Set("Authorization", token)

	if !isActive {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "токен недействителен"})
		c.Abort()
		return
	}

	c.Next()
}
