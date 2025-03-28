package postgres

import (
	"booking/models"
	"booking/pkg/helper/helper"
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func (u *userRepo) Create(ctx context.Context, req *models.CreateUser) (*models.UserPrimaryKey, error) {

	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	query := `INSERT INTO users(
		id,
		first_name,
		last_name,
		username,
		email,
		password_hash,
		role_id,
		created_at,
		updated_at 
	) VALUES ($1, $2, $3, $4, $5, $6, $7, now(), now())`

	_, err = u.db.Exec(ctx, query,
		uuid.String(),
		req.FirstName,
		req.LastName,
		req.UserName,
		req.Email,
		req.Password,
		req.RoleId,
	)

	id := uuid.String()
	pKey := &models.UserPrimaryKey{
		Id: &id,
	}

	return pKey, err
}

func (u *userRepo) GetByID(ctx context.Context, req *models.UserPrimaryKey) (*models.User, error) {
	res := &models.User{}
	query := `
		SELECT
			id,
			first_name,
			last_name,
			username,
			email,
			password_hash,
			role_id,
			created_at,
			updated_at
		FROM
			"users"
		WHERE
			id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.UserName,
		&res.Email,
		&res.Password,
		&res.RoleId,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *userRepo) GetList(ctx context.Context, req *models.GetListUserRequest) (*models.GetListUserResponse, error) {
	res := &models.GetListUserResponse{}
	params := make(map[string]interface{})
	var arr []interface{}
	query := `SELECT
		id,
		first_name,
		last_name,
		username,
		email,
		password_hash,
		role_id,
		created_at,
		updated_at
	FROM
		"users"`
	filter := " WHERE 1=1"
	order := " ORDER BY created_at"
	arrangement := " DESC"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND ((first_name || last_name || username || email) ILIKE ('%' || :search || '%'))"
	}

	if req.Offset > 0 {
		params["offset"] = req.Offset
		offset = " OFFSET :offset"
	}

	if req.Limit > 0 {
		params["limit"] = req.Limit
		limit = " LIMIT :limit"
	}

	cQ := `SELECT count(1) FROM "users"` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	q := query + filter + order + arrangement + offset + limit

	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.User{}

		err = rows.Scan(
			&obj.Id,
			&obj.FirstName,
			&obj.LastName,
			&obj.UserName,
			&obj.Email,
			&obj.Password,
			&obj.RoleId,
			&obj.CreatedAt,
			&obj.UpdatedAt,
		)

		if err != nil {
			return res, err
		}

		res.Users = append(res.Users, obj)
	}

	return res, nil
}

func (u *userRepo) Update(ctx context.Context, req *models.UpdateUser) (int64, error) {
	query := `UPDATE "users" SET `
	params := []interface{}{}
	counter := 1

	updated := false

	if req.FirstName != nil {
		query += `first_name = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.FirstName)
		counter++
		updated = true
	}

	if req.LastName != nil {
		query += `last_name = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.LastName)
		counter++
		updated = true
	}

	if req.UserName != nil {
		query += `username = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.UserName)
		counter++
		updated = true
	}

	if req.Email != nil {
		query += `email = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.Email)
		counter++
		updated = true
	}

	if req.Password != nil {
		query += `password_hash = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.Password)
		counter++
		updated = true
	}

	if req.RoleId != nil {
		query += `role_id = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.RoleId)
		counter++
		updated = true
	}

	if !updated {
		return 0, fmt.Errorf("no fields to update")
	}

	query = query[:len(query)-2] + `, updated_at = now()`
	query += ` WHERE id = $` + fmt.Sprint(counter)
	params = append(params, req.Id)

	result, err := u.db.Exec(ctx, query, params...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, nil
}

func (u *userRepo) Delete(ctx context.Context, req *models.UserPrimaryKey) (id int64, err error) {
	query := `DELETE FROM "users" WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

func (u *userRepo) GetByUserName(ctx context.Context, login *models.Login) (*models.User, error) {
	res := &models.User{}
	query := `
		SELECT
			id,
			first_name,
			last_name,
			username,
			email,
			password_hash,
			role_id,
			created_at,
			updated_at
		FROM
			"users"
		WHERE
			username = $1`

	err := u.db.QueryRow(ctx, query, login.UserName).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.UserName,
		&res.Email,
		&res.Password,
		&res.RoleId,
		&res.CreatedAt,
		&res.UpdatedAt,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}
