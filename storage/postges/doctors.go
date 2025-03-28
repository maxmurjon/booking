package postgres

import (
	"booking/models"
	"booking/pkg/helper/helper"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type doctorRepo struct {
	db *pgxpool.Pool
}

func (d *doctorRepo) Create(ctx context.Context, req *models.CreateDoctor) (*models.Doctor, error) {
	newUUID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO doctors(
		id,
		user_id,
		specialty,
		work_start,
		work_end,
		created_at
	) VALUES ($1, $2, $3, $4, $5, now())`

	_, err = d.db.Exec(ctx, query,
		newUUID.String(),
		req.UserId,
		req.Specialty,
		req.WorkStart,
		req.WorkEnd,
	)
	if err != nil {
		return nil, err
	}

	return &models.Doctor{
		Id:        newUUID.String(),
		UserId:    req.UserId,
		Specialty: req.Specialty,
		WorkStart: req.WorkStart,
		WorkEnd:   req.WorkEnd,
		CreatedAt: time.Now(),
	}, nil
}

func (d *doctorRepo) GetByID(ctx context.Context, id string) (*models.Doctor, error) {
	res := &models.Doctor{}
	query := `
		SELECT
			id,
			user_id,
			specialty,
			work_start,
			work_end,
			created_at
		FROM
			"doctors"
		WHERE
			id = $1`

	err := d.db.QueryRow(ctx, query, id).Scan(
		&res.Id,
		&res.UserId,
		&res.Specialty,
		&res.WorkStart,
		&res.WorkEnd,
		&res.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (d *doctorRepo) GetList(ctx context.Context, req *models.GetListDoctorRequest) (*models.GetListDoctorResponse, error) {
	res := &models.GetListDoctorResponse{}
	params := map[string]interface{}{}
	var arr []interface{}
	query := `SELECT
		id,
		user_id,
		specialty,
		work_start,
		work_end,
		created_at
	FROM
		"doctors"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND (specialty ILIKE ('%' || :search || '%'))"
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	// Count query
	cQ := `SELECT count(1) FROM "doctors"` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := d.db.QueryRow(ctx, cQ, arr...).Scan(&res.Count)
	if err != nil {
		return res, err
	}

	// Main query
	q := query + filter + order + arrangement + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)

	rows, err := d.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.Doctor{}
		err = rows.Scan(
			&obj.Id,
			&obj.UserId,
			&obj.Specialty,
			&obj.WorkStart,
			&obj.WorkEnd,
			&obj.CreatedAt,
		)
		if err != nil {
			return res, err
		}
		res.Doctors = append(res.Doctors, obj)
	}

	return res, nil
}

func (d *doctorRepo) Update(ctx context.Context, req *models.UpdateDoctor) (int64, error) {
	query := `UPDATE "doctors" SET `
	params := []interface{}{}
	counter := 1
	updated := false

	if req.UserId != "" {
		query += fmt.Sprintf(`user_id = $%d, `, counter)
		params = append(params, req.UserId)
		counter++
		updated = true
	}

	if req.Specialty != "" {
		query += fmt.Sprintf(`specialty = $%d, `, counter)
		params = append(params, req.Specialty)
		counter++
		updated = true
	}

	if !req.WorkStart.IsZero() {
		query += fmt.Sprintf(`work_start = $%d, `, counter)
		params = append(params, req.WorkStart)
		counter++
		updated = true
	}

	if !req.WorkEnd.IsZero() {
		query += fmt.Sprintf(`work_end = $%d, `, counter)
		params = append(params, req.WorkEnd)
		counter++
		updated = true
	}

	if !updated {
		return 0, fmt.Errorf("no fields to update")
	}

	query = query[:len(query)-2] + fmt.Sprintf(` WHERE id = $%d`, counter)
	params = append(params, req.Id)

	result, err := d.db.Exec(ctx, query, params...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()
	return rowsAffected, nil
}

func (d *doctorRepo) Delete(ctx context.Context, id string) (int64, error) {
	query := `DELETE FROM "doctors" WHERE id = $1`
	result, err := d.db.Exec(ctx, query, id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()
	return rowsAffected, nil
}
