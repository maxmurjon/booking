package handler

import (
	"booking/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateDoctor(c *gin.Context) {
	var entity *models.CreateDoctor
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	role,err:=h.strg.Role().GetByName(context.Background(),&models.Role{Name: "doctor"})
	if err!=nil{
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to get Doctor role: " + err.Error(),
		})
		return	
	}

	_,err=h.strg.User().Update(context.Background(),&models.UpdateUser{
		Id: entity.UserId,
		RoleId: &role.Id,
	})
	if err!=nil{
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

func (h *Handler) UpdateDoctor(c *gin.Context) {
	var entity models.UpdateDoctor
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Yaroqsiz so'rov tanasi: " + err.Error(),
		})
		return
	}

	if entity.Id == "" {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Foydalanuvchi ID-si talab qilinadi",
		})
		return
	}

	rowsAffected, err := h.strg.Doctor().Update(context.Background(), &entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Foydalanuvchini yangilashda xato: " + err.Error(),
		})
		return
	}

	if rowsAffected == 0 {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Foydalanuvchi topilmadi",
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Foydalanuvchi yangilandi",
		Data:    entity.Id,
	})
}

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