package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
)

// @Summary createPerm
// @Security ApiKeyAuth
// @Tags permission
// @Description create new permission(only for Superadmin)
// @Accept  json
// @Produce  json
// @Param input body model.CreatePerm true "Perm"
// @Success 201 {object} map[string]int
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /perms/ [post]
func (h *Handler) createPerm(ctx *gin.Context) {
	necessaryRole := []string{"Superadmin"}
	if err := h.services.CheckRole(necessaryRole, ctx.GetString("role")); err != nil {
		h.logger.Warnf("Handler createPerm:not enough rights")
		ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "not enough rights"})
		return
	}
	var input model.CreatePerm
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createPerm (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	permId, err := h.services.AuthUser.CreatePermission(input.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, map[string]int{
		"id": permId,
	})
}

// @Summary getAllPerms
// @Security ApiKeyAuth
// @Tags permission
// @Description gets all permissions(only for Superadmin)
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ListPerms
// @Failure 500 {object} model.ErrorResponse
// @Router /perms/ [get]
func (h *Handler) getAllPerms(ctx *gin.Context) {
	necessaryRole := []string{"Superadmin"}
	if err := h.services.CheckRole(necessaryRole, ctx.GetString("role")); err != nil {
		h.logger.Warnf("Handler getAllPerms:not enough rights")
		ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "not enough rights"})
		return
	}
	perms, err := h.services.AuthUser.GetAllPerms()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.ListPerms{Perms: perms})
}
