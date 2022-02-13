package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)


// @Summary Get Role By Id
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id
// @ID get-role-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object}
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