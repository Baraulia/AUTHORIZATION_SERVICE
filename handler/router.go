package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "stlab.itechart-group.com/go/food_delivery/authorization_service/docs"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/service"
)

type Handler struct {
	services *service.Service
	logger   logging.Logger
}

func NewHandler(services *service.Service, logger logging.Logger) *Handler {
	return &Handler{services: services, logger: logger}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Use(
		h.CorsMiddleware,
	)
	router.GET("/refresh", h.refreshToken)

	role := router.Group("/roles")
	{
		role.POST("/", h.createRole)
		role.GET("/:id", h.getRoleById)
		role.GET("/", h.getAllRoles)
		role.GET("/:id/perms", h.getPermsByRoleId)
		role.POST("/:id/perms", h.bindRoleWithPerms)
	}
	perm := router.Group("/perms")
	{
		perm.POST("/", h.createPerm)
		perm.GET("/", h.getAllPerms)
	}

	return router
}
