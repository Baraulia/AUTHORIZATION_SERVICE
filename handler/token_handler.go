package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
)

// @Summary refreshToken
// @Tags refresh
// @Description regeneration tokens by refresh token
// @Accept  json
// @Produce  json
// @Param refresh_token header string true "Refresh token"
// @Success 200 {object} authProto.GeneratedTokens
// @Failure 401 {object} model.ErrorResponse
// @Router /refresh [get]
func (h *Handler) refreshToken(ctx *gin.Context) {
	header := ctx.GetHeader("Refresh_token")
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
