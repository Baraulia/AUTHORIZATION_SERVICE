package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"strconv"
)

// @Summary getRoleById
// @Security ApiKeyAuth
// @Tags roles
// @Description get role by id
// @Accept  json
// @Produce  json
// @Param id path int true "RoleID"
// @Success 200 {object} model.Role
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /roles/{id} [get]
func (h *Handler) getRoleById(ctx *gin.Context) {
	necessaryRole := "Superadmin"
	if err := h.services.CheckRoleRights(nil, necessaryRole, ctx.GetString("perms"), ctx.GetString("role")); err != nil {
		h.logger.Warnf("Handler getRoleById:not enough rights")
		ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "not enough rights"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.logger.Warnf("Handler getRoleById (reading param):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	role, err := h.services.RolePerm.GetRoleById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, role)
}

// @Summary createRole
// @Security ApiKeyAuth
// @Tags roles
// @Description create new role
// @Accept  json
// @Produce  json
// @Param input body model.CreateRole true "Role"
// @Success 200 {object} map[string]int
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /roles/ [post]
func (h *Handler) createRole(ctx *gin.Context) {
	necessaryRole := "Superadmin"
	if err := h.services.CheckRoleRights(nil, necessaryRole, ctx.GetString("perms"), ctx.GetString("role")); err != nil {
		h.logger.Warnf("Handler getRoleById:not enough rights")
		ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "not enough rights"})
		return
	}
	var input model.CreateRole
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createRole (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	roleId, err := h.services.RolePerm.CreateRole(input.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, map[string]int{
		"id": roleId,
	})
}

// @Summary bindRoleWithPerms
// @Security ApiKeyAuth
// @Tags roles
// @Description binding role with permissions
// @Accept  json
// @Produce  json
// @Param input body model.BindRoleWithPermission true "Role and Perms"
// @Success 201
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /roles/{id}/perms [post]
func (h *Handler) bindRoleWithPerms(ctx *gin.Context) {
	necessaryRole := "Superadmin"
	if err := h.services.CheckRoleRights(nil, necessaryRole, ctx.GetString("perms"), ctx.GetString("role")); err != nil {
		h.logger.Warnf("Handler getRoleById:not enough rights")
		ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "not enough rights"})
		return
	}
	var input model.BindRoleWithPermission
	if err := ctx.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler BindRoleWithPerms (binding JSON):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	err := h.services.RolePerm.BindRoleWithPerms(&input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.AbortWithStatus(http.StatusCreated)
}

// @Summary getAllRoles
// @Security ApiKeyAuth
// @Tags roles
// @Description gets all roles
// @Accept  json
// @Produce  json
// @Success 200 {object} model.ListRoles
// @Failure 500 {object} model.ErrorResponse
// @Router /roles/ [get]
func (h *Handler) getAllRoles(ctx *gin.Context) {
	necessaryRole := "Superadmin"
	if err := h.services.CheckRoleRights(nil, necessaryRole, ctx.GetString("perms"), ctx.GetString("role")); err != nil {
		h.logger.Warnf("Handler getRoleById:not enough rights")
		ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "not enough rights"})
		return
	}
	roles, err := h.services.RolePerm.GetAllRoles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &model.ListRoles{Roles: roles})
}

// @Summary getPermsByRoleId
// @Security ApiKeyAuth
// @Tags roles
// @Description get permissions bound with role
// @Accept  json
// @Produce  json
// @Param id path int true "Role ID" Format(int64)
// @Success 200 {object} model.ListPerms
// @Failure 400 {object} model.ErrorResponse
// @Failure 500 {object} model.ErrorResponse
// @Router /roles/{id}/perms [get]
func (h *Handler) getPermsByRoleId(ctx *gin.Context) {
	necessaryRole := "Superadmin"
	if err := h.services.CheckRoleRights(nil, necessaryRole, ctx.GetString("perms"), ctx.GetString("role")); err != nil {
		h.logger.Warnf("Handler getRoleById:not enough rights")
		ctx.JSON(http.StatusUnauthorized, model.ErrorResponse{Message: "not enough rights"})
		return
	}
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		h.logger.Warnf("Handler getPermsByRoleId (reading param):%s", err)
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{Message: "invalid request"})
		return
	}
	perms, err := h.services.RolePerm.GetPermsByRoleId(id)
	if err != nil {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{Message: err.Error()})
			return
		}
	}
	ctx.JSON(http.StatusOK, &model.ListPerms{Perms: perms})
}
