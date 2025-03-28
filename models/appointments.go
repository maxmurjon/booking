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
    DoctorID       string    `json:"doctor_id,omitempty"`        
    PatientID      string    `json:"patient_id,omitempty"`       
    AppointmentDate time.Time `json:"appointment_date,omitempty"`
    Offset         int       `json:"offset"`                    
    Limit          int       `json:"limit"`                     
}


type GetListAppointmentResponse struct {
	Count        int            `json:"count"`
	Appointments []*Appointment `json:"Appointments"`
}
