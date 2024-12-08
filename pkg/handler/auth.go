package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vertexUP/models"
	"vertexUP/pkg/utils"
)

func (h Handler) signUp(c *gin.Context) {
	var input *models.SignUpInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, processStatus := h.usecase.SignUp(input)
	if processStatus != utils.Success {
		c.JSON(http.StatusBadRequest, processStatus)
		return
	}

	token, processStatus := h.usecase.SetToken(output.Login)
	if processStatus != utils.Success {
		c.JSON(http.StatusBadRequest, processStatus)
		return
	}

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, output)
}

func (h Handler) signIn(c *gin.Context) {
	var input *models.SignInInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, processStatus := h.usecase.SignIn(input)
	if processStatus != utils.Success {
		c.JSON(http.StatusBadRequest, processStatus)
		return
	}

	token, processStatus := h.usecase.SetToken(output.Login)
	if processStatus != utils.Success {
		c.JSON(http.StatusBadRequest, processStatus)
		return
	}

	c.Header("Authorization", token)

	c.JSON(http.StatusOK, output)
}

func (h Handler) logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "токен отсутствует"})
		return
	}

	processStatus := h.usecase.Logout(token)
	if processStatus != utils.Success {
		c.JSON(http.StatusUnauthorized, processStatus)
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
