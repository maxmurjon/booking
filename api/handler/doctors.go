package handler

import (
	"booking/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateDoctor godoc
// @Summary Create a new doctor
// @Description Assigns a user the "doctor" role and creates a doctor profile
// @Tags doctors
// @Accept json
// @Produce json
// @Param doctor body models.CreateDoctor true "Doctor Data"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.DefaultError
// @Router /doctors [post]
func (h *Handler) CreateDoctor(c *gin.Context) {
	var entity *models.CreateDoctor
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	
	role, err := h.strg.Role().GetByName(context.Background(), &models.Role{Name: "doctor"})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to get Doctor role: " + err.Error(),
		})
		return
	}

	_, err = h.strg.User().Update(context.Background(), &models.UpdateUser{
		Id:     entity.UserId,
		RoleId: &role.Id,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to change User's Role to Doctor: " + err.Error(),
		})
		return
	}

	Doctor, err := h.strg.Doctor().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Doctor: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Doctor has been created",
		Data:    Doctor,
	})
}

// UpdateDoctor godoc
// @Summary Update doctor information
// @Description Updates details of an existing doctor
// @Tags doctors
// @Accept json
// @Produce json
// @Param doctor body models.UpdateDoctor true "Doctor Data"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.DefaultError
// @Failure 404 {object} models.DefaultError
// @Router /doctors [put]
func (h *Handler) UpdateDoctor(c *gin.Context) {
	var entity models.UpdateDoctor
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	if entity.Id == "" {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Doctor ID is required",
		})
		return
	}

	rowsAffected, err := h.strg.Doctor().Update(context.Background(), &entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to update Doctor: " + err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Doctor not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Doctor has been updated",
		Data:    entity.Id,
	})
}

// GetDoctorsList godoc
// @Summary Get a list of doctors
// @Description Retrieves a list of all doctors
// @Tags doctors
// @Produce json
// @Success 200 {object} models.GetListDoctorResponse
// @Failure 500 {object} models.DefaultError
// @Router /doctors [get]
func (h *Handler) GetDoctorsList(c *gin.Context) {
	resp, err := h.strg.Doctor().GetList(context.Background(), &models.GetListDoctorRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Doctor list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

// GetDoctorByID godoc
// @Summary Get a doctor by ID
// @Description Retrieves details of a specific doctor by ID
// @Tags doctors
// @Produce json
// @Param id path string true "Doctor ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.DefaultError
// @Router /doctors/{id} [get]
func (h *Handler) GetDoctorsByIDHandler(c *gin.Context) {
	id := c.Param("id")

	Doctor, err := h.strg.Doctor().GetByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Doctor not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    Doctor,
	})
}

// DeleteDoctor godoc
// @Summary Delete a doctor
// @Description Deletes a doctor by ID
// @Tags doctors
// @Param id path string true "Doctor ID"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.DefaultError
// @Failure 500 {object} models.DefaultError
// @Router /doctors/{id} [delete]
func (h *Handler) DeleteDoctor(c *gin.Context) {
	id := c.Param("id")

	rowsAffected, err := h.strg.Doctor().Delete(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete Doctor: " + err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Doctor not found",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Doctor has been deleted",
		Data:    id,
	})
}