package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"strings"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		h.logger.Errorf("userIdentity:empty auth header")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: "empty auth header"})
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		h.logger.Errorf("userIdentity:invalid auth header")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: "invalid auth header"})
		return
	}
	if len(headerParts[1]) == 0 {
		h.logger.Errorf("userIdentity:token is empty")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: "token is empty"})
		return
	}
	ok, err := h.services.Authorization.CheckRights(headerParts[1], "Superadmin")
	if err != nil || !ok {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.Next()
}

func (h *Handler) CorsMiddleware(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Header("Access-Control-Allow-Methods", "*")
	ctx.Header("Access-Control-Allow-Headers", "*")
	ctx.Header("Content-Type", "application/json")

	if ctx.Request.Method != "OPTIONS" {
		ctx.Next()
	} else {
		ctx.AbortWithStatus(http.StatusOK)
	}
}
