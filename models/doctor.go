package models

import "time"

type Doctor struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Specialty string    `json:"specialty"`
	WorkStart time.Time `json:"work_start"` // Vaqt turi bilan almashtirildi
	WorkEnd   time.Time `json:"work_end"`   // Vaqt turi bilan almashtirildi
	CreatedAt time.Time `json:"created_at"`
}

type CreateDoctor struct {
	UserId    string    `json:"user_id"`
	Specialty string    `json:"specialty"`
	WorkStart time.Time `json:"work_start"`
	WorkEnd   time.Time `json:"work_end"`
}

type UpdateDoctor struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	Specialty string    `json:"specialty"`
	WorkStart time.Time `json:"work_start"`
	WorkEnd   time.Time `json:"work_end"`
}

type GetListDoctorRequest struct {
	Offset int64  `json:"offset"` // Sahifalash uchun int64 
	Limit  int64  `json:"limit"`  // Sahifalash uchun int64
	Search string `json:"search"`
}

type GetListDoctorResponse struct {
	Count   int       `json:"count"`
	Doctors []*Doctor `json:"doctors"`
}
