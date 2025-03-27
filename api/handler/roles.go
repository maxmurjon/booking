package handler

import (
	"comics/models"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRole(c *gin.Context) {
	var entity *models.CreateRole
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Create the Role in the storage
	id, err := h.strg.Role().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Role: " + err.Error(),
		})
		return
	}

	// Get the Role by ID
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

func (h *Handler) UpdateRole(c *gin.Context) {
	var entity models.UpdateRole
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	fmt.Println(entity)

	// Update the Role
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

func (h *Handler) GetRolesList(c *gin.Context) {
	// Retrieve the list of Roles (offset and limit can be implemented later)
	resp, err := h.strg.Role().GetList(context.Background(), &models.GetListRoleRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Role list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetRolesByIDHandler(c *gin.Context) {
	id := c.Param("id")
	intId,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Role not found: " + err.Error(),
		})
		return
	}
	// Get the Role by ID
	role, err := h.strg.Role().GetByID(context.Background(), &models.PrimaryKey{Id: intId})
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

func (h *Handler) DeleteRole(c *gin.Context) {
	id := c.Param("id")

	intId,err:=strconv.Atoi(id)
	if err!=nil{
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Role not found: " + err.Error(),
		})
		return
	}
	// Delete the Role by ID
	deletedRole, err := h.strg.Role().Delete(context.Background(), &models.PrimaryKey{Id: intId})
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
