package handler

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
	"vertexUP/models"
	"vertexUP/pkg/service"
	"vertexUP/pkg/usecase"
	"vertexUP/pkg/utils"
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
	allowOrigins := strings.Split(env.Domain, ",")

	router.Use(cors.New(cors.Config{
		AllowOrigins: allowOrigins,
		AllowMethods: []string{http.MethodPut, http.MethodPost, http.MethodGet, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{"Content-Status", "Access-Control-Allow-Headers", "Access-Control-Allow-Origin",
			utils.HeaderAuthorization, utils.HeaderClientRequestId},
		ExposeHeaders: []string{"Content-Length", utils.HeaderTimestamp,
			utils.HeaderClientRequestId, utils.HeaderRequestId},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	auth := router.Group("api/auth")
	{
		auth.POST("/sign-up", h.signUp)
	}

	return router
}

func ginSetMode(serverMode string) {
	if serverMode == ServerMode.RELEASE {
		gin.SetMode(gin.ReleaseMode)
	}
}
