package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"strconv"
)

// @Summary createPerm
// @Tags permission
// @Description create new permission
// @Accept  json
// @Produce  json
// @Param input body model.CreatePerm true "Perm"
// @Success 201 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /perms/ [post]
func (h *Handler) createPerm(ctx *gin.Context) {
	var input model.CreatePerm
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createUser (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	permId, err := h.services.RolePerm.CreatePermission(input.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, map[string]int{
		"id": permId,
	})
}

// @Summary getPermsByRoleId
// @Tags permission
// @Description get permissions bound with role
// @Accept  json
// @Produce  json
// @Param id path int true "Role ID" Format(int64)
// @Success 200 {object} model.ListPerms
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /perms/{id} [get]
func (h *Handler) getPermsByRoleId(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.logger.Warnf("Handler getPermsByRoleId (reading param):%s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	perms, err := h.services.RolePerm.GetPermsByRoleId(id)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, &model.ListPerms{Perms: perms})
}

// @Summary getAllPerms
// @Tags permission
// @Description get all permissions
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ListPerms
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /perms/ [get]
func (h *Handler) getAllPerms(ctx *gin.Context) {
	perms, err := h.services.RolePerm.GetAllPerms()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.ListPerms{Perms: perms})
}
