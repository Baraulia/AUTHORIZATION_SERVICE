package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"strconv"
)

// @Summary getRoleById
// @Tags roles
// @Description get role by id
// @Accept  json
// @Produce  json
// @Param id path int true "RoleID"
// @Success 200 {object} model.Role
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /roles/{id} [get]
func (h *Handler) getRoleById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.logger.Warnf("Handler getUser (reading param):%s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	role, err := h.services.RolePerm.GetRoleById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, role)
}

// @Summary createRole
// @Tags roles
// @Description create new role
// @Accept  json
// @Produce  json
// @Param input body model.CreateRole true "Role"
// @Success 200 {string} string
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /roles/ [post]
func (h *Handler) createRole(ctx *gin.Context) {
	var input model.CreateRole
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createRole (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	roleId, err := h.services.RolePerm.CreateRole(input.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, map[string]int{
		"id": roleId,
	})
}

// @Summary bindRoleWithPerms
// @Tags roles
// @Description binding role with permissions
// @Accept  json
// @Produce  json
// @Param input body model.BindRoleWithPermission true "Role and Perms"
// @Success 201
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /roles/roleToPerms [post]
func (h *Handler) bindRoleWithPerms(ctx *gin.Context) {
	var input model.BindRoleWithPermission
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler BindRoleWithPerms (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	err := h.services.RolePerm.BindRoleWithPerms(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.AbortWithStatus(http.StatusCreated)
}

// @Summary getAllRoles
// @Tags roles
// @Description gets all roles
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ListRoles
// @Failure 400 {string} string
// @Failure 500 {string} string
// @Router /roles/ [get]
func (h *Handler) getAllRoles(ctx *gin.Context) {
	roles, err := h.services.RolePerm.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.ListRoles{Roles: roles})
}
