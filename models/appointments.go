package models

import (
	"time"
)

type Appointment struct {
	ID              int       `json:"id"`
	PatientID       string    `json:"patient_id"`
	DoctorID        string    `json:"doctor_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	Status          string    `json:"status"`
	CreatedAt       time.Time `json:"created_at"`
}

type CreateAppointment struct {
	ID              int       `json:"id"`
	PatientID       string    `json:"patient_id"`
	DoctorID        string    `json:"doctor_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
}

type UpdateAppointment struct {
	ID              int       `json:"id"`
	PatientID       string    `json:"patient_id"`
	DoctorID        string    `json:"doctor_id"`
	AppointmentDate time.Time `json:"appointment_date"`
	AppointmentTime time.Time `json:"appointment_time"`
	Status          string    `json:"status"`
}

type GetListAppointmentRequest struct {
	Offset int64  `json:"offset"` // Sahifalash uchun int64
	Limit  int64  `json:"limit"`  // Sahifalash uchun int64
	Search string `json:"search"`
}

type GetListAppointmentResponse struct {
	Count        int            `json:"count"`
	Appointments []*Appointment `json:"Appointments"`
}
