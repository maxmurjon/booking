package handler

import (
	"booking/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRole godoc
// @Summary Create a new role
// @Description Create a new role with given details
// @Tags roles
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param role body models.CreateRole true "Role data"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.DefaultError
// @Failure 500 {object} models.DefaultError
// @Router /roles [post]
func (h *Handler) CreateRole(c *gin.Context) {
	var entity *models.CreateRole
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	id, err := h.strg.Role().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Role: " + err.Error(),
		})
		return
	}

	role, err := h.strg.Role().GetByID(context.Background(), &models.PrimaryKey{Id: id.Id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created Role: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Role has been created",
		Data:    role,
	})
}

// UpdateRole godoc
// @Summary Update an existing role
// @Description Update role details
// @Tags roles
// @Accept json
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param role body models.UpdateRole true "Updated role data"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.DefaultError
// @Failure 500 {object} models.DefaultError
// @Router /roles [put]
func (h *Handler) UpdateRole(c *gin.Context) {
	var entity models.UpdateRole
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	if _, err := h.strg.Role().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update Role: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Role has been updated",
		Data:    "ok",
	})
}

// GetRolesList godoc
// @Summary Get list of roles
// @Description Retrieve all roles
// @Tags roles
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Success 200 {object} models.GetListRoleResponse
// @Failure 500 {object} models.DefaultError
// @Router /roles [get]
func (h *Handler) GetRolesList(c *gin.Context) {
	resp, err := h.strg.Role().GetList(context.Background(), &models.GetListRoleRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Role list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetRolesByIDHandler godoc
// @Summary Get role by ID
// @Description Retrieve a role using ID
// @Tags roles
// @Produce json
// @Param Authorization header string true "Authorization token"
// @Param id path string true "Role ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.DefaultError
// @Router /roles/{id} [get]
func (h *Handler) GetRolesByIDHandler(c *gin.Context) {
	id := c.Param("id")

	role, err := h.strg.Role().GetByID(context.Background(), &models.PrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Role not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    role,
	})
}

// DeleteRole godoc
// @Summary Delete role by ID
// @Description Delete a role using ID
// @Tags roles
// @Param Authorization header string true "Authorization token"
// @Param id path string true "Role ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.DefaultError
// @Router /roles/{id} [delete]
func (h *Handler) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	deletedRole, err := h.strg.Role().Delete(context.Background(), &models.PrimaryKey{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete Role: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Role has been deleted",
		Data:    deletedRole,
	})
}