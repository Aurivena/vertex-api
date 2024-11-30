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

	c.JSON(http.StatusOK, output)
}
