package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
)

func (h *Handler) refreshToken(ctx *gin.Context) {
	header := ctx.GetHeader("Refresh")
	if header == "" {
		h.logger.Errorf("refreshToken:empty refresh header")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: "empty refresh header"})
		return
	}
	tokens, err := h.services.Authorization.RefreshTokens(header)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tokens)
}
