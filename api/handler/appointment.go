package handler

import (
	"booking/models"
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateAppointment(c *gin.Context) {
	var entity models.CreateAppointment
	if err := c.BindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Invalid request body: " + err.Error(),
		})
		return
	}

	doctor, err := h.strg.Doctor().GetByID(c.Request.Context(), entity.DoctorID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Doctor not found: " + err.Error(),
		})
		return
	}

	doctorStart := time.Date(entity.AppointmentDate.Year(), entity.AppointmentDate.Month(), entity.AppointmentDate.Day(),
		doctor.WorkStart.Hour(), doctor.WorkStart.Minute(), 0, 0, entity.AppointmentDate.Location())

	doctorEnd := time.Date(entity.AppointmentDate.Year(), entity.AppointmentDate.Month(), entity.AppointmentDate.Day(),
		doctor.WorkEnd.Hour(), doctor.WorkEnd.Minute(), 0, 0, entity.AppointmentDate.Location())

	appointmentTime := time.Date(entity.AppointmentDate.Year(), entity.AppointmentDate.Month(), entity.AppointmentDate.Day(),
		entity.AppointmentTime.Hour(), entity.AppointmentTime.Minute(), 0, 0, entity.AppointmentTime.Location())

	if appointmentTime.Before(doctorStart) || appointmentTime.After(doctorEnd) {
		log.Println("Doctor not working at this time")
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: fmt.Sprintf("Doctor works between %v and %v", doctor.WorkStart, doctor.WorkEnd),
		})
		return
	}

	doctorAppointments, err := h.strg.Appointment().GetList(c.Request.Context(), &models.GetListAppointmentRequest{
		DoctorID:        entity.DoctorID,
		AppointmentDate: entity.AppointmentDate,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve Appointment list: " + err.Error(),
		})
		return
	}
	
	for _, appointment := range doctorAppointments.Appointments {
		appointmentEndTime := appointment.AppointmentTime.Add(time.Hour)
		appointmentStartTime := appointment.AppointmentTime.Add(-time.Hour)

		if entity.AppointmentTime.Before(appointmentEndTime) && entity.AppointmentTime.After(appointmentStartTime) {
			log.Println("Doctor has an appointment at this time")
			c.JSON(http.StatusBadRequest, models.DefaultError{
				Message: "Doctor has an appointment at this time",
			})
			return
		}
	}

	id, err := h.strg.Appointment().Create(c.Request.Context(), &entity)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.DefaultError{
			Message: "Failed to create Appointment: " + err.Error(),
		})
		return
	}

	appointment, err := h.strg.Appointment().GetByID(c.Request.Context(), id.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.DefaultError{
			Message: "Failed to retrieve created Appointment: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.SuccessResponse{
		Message: "Appointment has been created",
		Data:    appointment,
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
	intId, err := strconv.Atoi(id)
	if err != nil {
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
	intId, err := strconv.Atoi(id)
	if err != nil {
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
