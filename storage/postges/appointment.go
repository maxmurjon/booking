package postgres

import (
	"booking/models"
	"context"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type appointmentRepo struct {
	db *pgxpool.Pool
}

// Create appointment
func (r *appointmentRepo) Create(ctx context.Context, req *models.CreateAppointment) (*models.Appointment, error) {
	query := `INSERT INTO appointments (
		patient_id, doctor_id, appointment_date, appointment_time, status, created_at
	) VALUES ($1, $2, $3, $4, 'scheduled', now()) RETURNING id, created_at`

	var appointment models.Appointment
	err := r.db.QueryRow(ctx, query,
		req.PatientID, req.DoctorID, req.AppointmentDate, req.AppointmentTime,
	).Scan(&appointment.ID, &appointment.CreatedAt)
	if err != nil {
		return nil, err
	}

	appointment.PatientID = req.PatientID
	appointment.DoctorID = req.DoctorID
	appointment.AppointmentDate = req.AppointmentDate
	appointment.AppointmentTime = req.AppointmentTime
	appointment.Status = "scheduled"

	return &appointment, nil
}

// Get appointment by ID
func (r *appointmentRepo) GetByID(ctx context.Context, id int) (*models.Appointment, error) {
	query := `SELECT id, patient_id, doctor_id, appointment_date, appointment_time, status, created_at FROM appointments WHERE id = $1`
	appointment := &models.Appointment{}

	err := r.db.QueryRow(ctx, query, id).Scan(
		&appointment.ID,
		&appointment.PatientID,
		&appointment.DoctorID,
		&appointment.AppointmentDate,
		&appointment.AppointmentTime,
		&appointment.Status,
		&appointment.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return appointment, nil
}

// Get list of appointments
func (r *appointmentRepo) GetList(ctx context.Context, req *models.GetListAppointmentRequest) (*models.GetListAppointmentResponse, error) {
	response := &models.GetListAppointmentResponse{}

	// Asosiy so‘rov
	query := `SELECT id, patient_id, doctor_id, appointment_date, appointment_time, status, created_at 
			  FROM appointments`
	var args []interface{}
	var conditions []string

	// Search filter
	if req.Search != "" {
		conditions = append(conditions, `(CAST(patient_id AS TEXT) ILIKE $1 OR CAST(doctor_id AS TEXT) ILIKE $1 OR status ILIKE $1)`)
		args = append(args, "%"+req.Search+"%")
	}

	// WHERE shartini qo‘shish
	if len(conditions) > 0 {
		query += " WHERE " + strings.Join(conditions, " AND ")
	}

	// ORDER, OFFSET, LIMIT qo‘shish
	args = append(args, req.Offset, req.Limit)
	query += " ORDER BY created_at DESC OFFSET $" + strconv.Itoa(len(args)-1) + " LIMIT $" + strconv.Itoa(len(args))

	// So‘rovni bajarish
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		appointment := &models.Appointment{}
		if err := rows.Scan(
			&appointment.ID,
			&appointment.PatientID,
			&appointment.DoctorID,
			&appointment.AppointmentDate,
			&appointment.AppointmentTime,
			&appointment.Status,
			&appointment.CreatedAt,
		); err != nil {
			return nil, err
		}
		response.Appointments = append(response.Appointments, appointment)
	}

	return response, nil
}

// Update appointment
func (r *appointmentRepo) Update(ctx context.Context, req *models.UpdateAppointment) error {
	query := `UPDATE appointments SET patient_id = $1, doctor_id = $2, appointment_date = $3, appointment_time = $4, status = $5 WHERE id = $6`

	_, err := r.db.Exec(ctx, query,
		req.PatientID, req.DoctorID, req.AppointmentDate, req.AppointmentTime, req.Status, req.ID,
	)
	return err
}

// Delete appointment
func (r *appointmentRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM appointments WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
