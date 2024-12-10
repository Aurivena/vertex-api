package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"vertexUP/models"
	"vertexUP/pkg/utils"
)

// @Summary      Зарегистрировать пользователя
// @Description  Регистрирует пользователя в система и устанавливает ему jwt токен
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        models.SignUpInput body models.SignUpInput true "Входные данные"
// @Success      200 {object}  models.SignUpOutput  "Выходные данные"
// @Failure      400 {object} string "BadRequest"
// @Failure      500 {object} string "InternalServerError"
// @Router       /auth/sign-up [post]
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

// @Summary      Авторизовать в системе пользователя
// @Description  Авторизует пользователя в системе и выдает ему новый jwt токен
// @Tags         Account
// @Accept       json
// @Produce      json
// @Param        models.SignInInput body models.SignInInput true "Входные данные"
// @Success      200 {object}  models.SignInOutput   "Выходные данные"
// @Failure      400 {object} string "BadRequest"
// @Failure      400 {object} string "UnregisteredAccount"
// @Failure      500 {object} string "InternalServerError"
// @Router       /auth/sign-in [post]
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

// @Summary      Совершить выход из аккаунта
// @Description  Выходит из аккаунта и удаляет jwt токен
// @Tags         Account
// @Accept       json
// @Produce      json
// @Success      204 {object} string "NoContent"
// @Success      400 {object} string "BadRequest"
// @Failure      401 {object} string "StatusUnauthorized"
// @Failure      500 {object} string "InternalServerError"
// @Router       /account/logout [post]
func (h Handler) logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "токен отсутствует"})
		return
	}

	processStatus := h.usecase.Logout(token)
	if processStatus != utils.NoContent {
		c.JSON(http.StatusBadRequest, processStatus)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
