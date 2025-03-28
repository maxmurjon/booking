package postgres

import (
	"booking/models"
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type appointmentRepo struct {
	db *pgxpool.Pool
}

func (r *appointmentRepo) Create(ctx context.Context, req *models.CreateAppointment) (*models.Appointment, error) {
	fmt.Println(req)
	query := `INSERT INTO appointments (
		patient_id, doctor_id, appointment_date, appointment_time, status, created_at
	) VALUES ($1, $2, $3, $4, 'pending', now()) RETURNING id`

	var appointment models.Appointment
	err := r.db.QueryRow(ctx, query,
		req.PatientID, req.DoctorID, req.AppointmentDate, req.AppointmentTime,
	).Scan(&appointment.ID)
	if err != nil {
		return nil, err
	}

	appointment.PatientID = req.PatientID
	appointment.DoctorID = req.DoctorID
	appointment.AppointmentDate = req.AppointmentDate
	appointment.AppointmentTime = req.AppointmentTime
	appointment.Status = "pending"

	return &appointment, nil
}

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

func (r *appointmentRepo) GetList(ctx context.Context, req *models.GetListAppointmentRequest) (*models.GetListAppointmentResponse, error) {
	response := &models.GetListAppointmentResponse{}

	if req.Offset < 0 {
		req.Offset = 0
	}
	if req.Limit <= 0 {
		req.Limit = 10 
	}

	query := `SELECT id, patient_id, doctor_id, appointment_date, appointment_time, status, created_at 
	          FROM appointments WHERE 1=1`

	args := []interface{}{}
	argIndex := 1

	if req.DoctorID != "" {
		query += fmt.Sprintf(" AND doctor_id = $%d", argIndex)
		args = append(args, req.DoctorID)
		argIndex++
	}
	if !req.AppointmentDate.IsZero() {
		query += fmt.Sprintf(" AND appointment_date = $%d", argIndex)
		args = append(args, req.AppointmentDate)
		argIndex++
	}

	query += fmt.Sprintf(" ORDER BY created_at DESC OFFSET $%d LIMIT $%d", argIndex, argIndex+1)
	args = append(args, req.Offset, req.Limit)

	fmt.Println("Generated Query:", query)
	fmt.Println("Query Arguments:", args)

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		fmt.Println("Query Execution Error:", err)
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

	countQuery := `SELECT COUNT(*) FROM appointments WHERE 1=1`
	countArgs := []interface{}{}
	countIndex := 1

	if req.DoctorID != "" {
		countQuery += fmt.Sprintf(" AND doctor_id = $%d", countIndex)
		countArgs = append(countArgs, req.DoctorID)
		countIndex++
	}
	if !req.AppointmentDate.IsZero() {
		countQuery += fmt.Sprintf(" AND appointment_date = $%d", countIndex)
		countArgs = append(countArgs, req.AppointmentDate)
		countIndex++
	}

	fmt.Println("Count Query:", countQuery)
	fmt.Println("Count Arguments:", countArgs)

	err = r.db.QueryRow(ctx, countQuery, countArgs...).Scan(&response.Count)
	if err != nil {
		return nil, err
	}

	if len(response.Appointments) == 0 {
		fmt.Println("No appointments found with the given filters.")
	}

	return response, nil
}


func (r *appointmentRepo) Update(ctx context.Context, req *models.UpdateAppointment) error {
	query := `UPDATE appointments SET patient_id = $1, doctor_id = $2, appointment_date = $3, appointment_time = $4, status = $5 WHERE id = $6`

	_, err := r.db.Exec(ctx, query,
		req.PatientID, req.DoctorID, req.AppointmentDate, req.AppointmentTime, req.Status, req.ID,
	)
	return err
}

func (r *appointmentRepo) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM appointments WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
