package handler

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
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
			lists.GET("/:id", h.getRoleById)
		}
	}

	return router
}
