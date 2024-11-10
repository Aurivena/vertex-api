package handler

import (
	"github.com/gin-gonic/gin"
	"vertexUP/models"
	"vertexUP/pkg/service"
	"vertexUP/pkg/usecase"
	"vertexUP/server/ServerMode"
)

type Handler struct {
	usecase *usecase.Usecase
	service *service.Service
}

func NewHandler(usecase *usecase.Usecase, service *service.Service) *Handler {
	return &Handler{usecase: usecase, service: service}
}

func (h *Handler) InitHTTPRoutes(env *models.Environment) *gin.Engine {
	ginSetMode(env.ServerMode)
	router := gin.Default()

	return router
}

func ginSetMode(serverMode string) {
	if serverMode == ServerMode.RELEASE {
		gin.SetMode(gin.ReleaseMode)
	}
}
