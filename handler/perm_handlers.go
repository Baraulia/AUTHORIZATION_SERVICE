package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
)

// @Summary createPerm

// @Tags permission
// @Description create new permission
// @Accept  json
// @Produce  json
// @Param input body model.CreatePerm true "Perm"
// @Success 201 {object} map[string]int
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /perms/ [post]
func (h *Handler) createPerm(ctx *gin.Context) {
	var input model.CreatePerm
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createUser (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	permId, err := h.services.RolePerm.CreatePermission(input.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, map[string]int{
		"id": permId,
	})
}

// @Summary getAllPerms

// @Tags permission
// @Description gets all permissions
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ListPerms
// @Failure 500 {object} model.ErrorResponse
// @Router /perms/ [get]
func (h *Handler) getAllPerms(ctx *gin.Context) {
	perms, err := h.services.RolePerm.GetAllPerms()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.ListPerms{Perms: perms})
}
