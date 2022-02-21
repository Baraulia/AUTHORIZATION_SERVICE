package handler

import (
	"github.com/Baraulia/AUTHORIZATION_SERVICE/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// @Summary Get Role By Id
// @Security KeyAuth
// @Tags Role
// @Description get role by id
// @ID get-role-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Roles
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/roles/:id [get]
func (h *Handler) getRoleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	list, err := h.services.RoleList.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, list)
}

// @Summary Create Role
// @Security KeyAuth
// @Tags Role
// @Description create role
// @ID create-role
// @Accept  json
// @Produce  json
// @Param input body model.Role true "role info"
// @Success 200 {object} model.Role
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/roles/ [post]
func (h *Handler) createRole(c *gin.Context) {
	var input model.Role
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createUser (binding JSON):%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	role, err := h.services.RoleList.CreateRole(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, role)
}

// @Summary Create Permission
// @Security KeyAuth
// @Tags Permission
// @Description create permission
// @ID create-permission
// @Accept  json
// @Produce  json
// @Param input body model.Permission true "permission info"
// @Success 200 {object} model.Permission
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/roles/permission [post]
func (h *Handler) createPermission(c *gin.Context) {
	var input model.Permission
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createUser (binding JSON):%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	permission, err := h.services.RoleList.CreatePermission(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, permission)
}

// @Summary Create Role to Permission
// @Security KeyAuth
// @Tags Permission
// @Description create role to permission
// @ID create-role-to-permission
// @Accept  json
// @Produce  json
// @Param input body model.RoleToPermission true "RTP info"
// @Success 200 {object} model.RoleToPermission
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/roles/roleToPermission [post]
func (h *Handler) createRoleToPermission(c *gin.Context) {
	var input model.RoleToPermission
	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.Warnf("Handler createUser (binding JSON):%s", err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}
	permission, err := h.services.RoleList.CreateRoleToPermission(&input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, permission)
}