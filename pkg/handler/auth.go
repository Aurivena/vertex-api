package handler

import (
	"github.com/gin-gonic/gin"
	"vertexUP/models"
	"vertexUP/pkg/usecase"
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
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	output, processStatus := h.usecase.SignUp(input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	token, processStatus := h.usecase.SetToken(output.Login)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	c.Header("Authorization", token)

	h.sendResponseSuccess(c, &output, processStatus)
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
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	output, processStatus := h.usecase.SignIn(input)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	token, processStatus := h.usecase.SetToken(output.Login)
	if processStatus != usecase.Success {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	c.Header("Authorization", token)

	h.sendResponseSuccess(c, &output, processStatus)
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
// @Router       /account/logout [delete]
func (h Handler) logout(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		h.sendResponseSuccess(c, nil, usecase.BadRequest)
		return
	}

	processStatus := h.usecase.Logout(token)
	if processStatus != usecase.NoContent {
		h.sendResponseSuccess(c, nil, processStatus)
		return
	}

	h.sendResponseSuccess(c, nil, processStatus)
}
