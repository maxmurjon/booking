package handler

import (
	"booking/models"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAppointment(c *gin.Context) {
	var entity *models.CreateAppointment
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}
	doctor,err:=h.strg.Doctor().GetByID(context.Background(), entity.DoctorID)
	if err!=nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Doctor not found: " + err.Error(),
		})
		return
	}

	if doctor.WorkStart.Before(entity.AppointmentTime) || doctor.WorkEnd.After(entity.AppointmentTime) {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Doctor not working at this time",
		})
		return
	}

	

	

	id, err := h.strg.Appointment().Create(context.Background(), entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Appointment: " + err.Error(),
		})
		return
	}

	Appointment, err := h.strg.Appointment().GetByID(context.Background(), id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created Appointment: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Appointment has been created",
		Data:    Appointment,
	})
}

func (h *Handler) UpdateAppointment(c *gin.Context) {
	var entity models.UpdateAppointment
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Yaroqsiz so'rov tanasi: " + err.Error(),
		})
		return
	}

	if entity.ID == 0 {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Foydalanuvchi ID-si talab qilinadi",
		})
		return
	}

	if err := h.strg.Appointment().Update(context.Background(), &entity); err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Foydalanuvchini yangilashda xato: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Foydalanuvchi yangilandi",
		Data:    entity.ID,
	})
}

func (h *Handler) GetAppointmentsList(c *gin.Context) {
	resp, err := h.strg.Appointment().GetList(context.Background(), &models.GetListAppointmentRequest{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Appointment list: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetAppointmentsByIDHandler(c *gin.Context) {
	id := c.Param("id")
	intId ,err :=strconv.Atoi(id)
	if err!=nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Appointment id invalid: " + err.Error(),
		})
		return
	}

	Appointment, err := h.strg.Appointment().GetByID(context.Background(), intId)
	if err != nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Appointment not found: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "OK",
		Data:    Appointment,
	})
}

func (h *Handler) DeleteAppointment(c *gin.Context) {
	id := c.Param("id")
	intId ,err :=strconv.Atoi(id)
	if err!=nil {
		c.JSON(http.StatusNotFound, models.DefaultError{
			Message: "Appointment id invalid: " + err.Error(),
		})
		return
	}

	err = h.strg.Appointment().Delete(context.Background(), intId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to delete Appointment: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Appointment has been deleted",
		Data:    id,
	})
}
