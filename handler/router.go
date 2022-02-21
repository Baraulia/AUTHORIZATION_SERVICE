package handler

import (
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"

	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	_ "github.com/Baraulia/AUTHORIZATION_SERVICE/docs"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	logger  logging.Logger
}

func NewHandler(services *service.Service, logger logging.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(
		CorsMiddleware,
	)

	auth := router.Group("/auth")
	{
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/roles")
		{
			lists.POST("/", h.createRole)
			lists.POST("/permission", h.createPermission)
			lists.POST("/roleToPermission", h.createRoleToPermission)
			lists.GET("/:id", h.getRoleById)
		}
	}

	return router
}
