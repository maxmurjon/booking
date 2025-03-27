package postgres

import (
	"comics/models"
	"comics/pkg/helper/helper"
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
		password_hash,
		phone_number,
		created_at,
		updated_at 
	) VALUES ($1, $2, $3, $4, $5, now(), now())`

	_, err = u.db.Exec(ctx, query,
		uuid.String(),
		req.FirstName,
		req.LastName,
		req.Password,
		req.PhoneNumber,
	)

	pKey := &models.UserPrimaryKey{
		Id: uuid.String(),
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
            password_hash,
            phone_number,
			image_url,
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
		&res.Password,
		&res.PhoneNumber,
		&res.ImageUrl,
		&res.CreatedAt, // created_at as time.Time
		&res.UpdatedAt, // updated_at as time.Time
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
		password_hash,
		phone_number,
		image_url,
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
		filter += " AND ((name || phone || is_active || login) ILIKE ('%' || :search || '%'))"
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

		// Ensure ImageUrl is handled as a pointer to string
		err = rows.Scan(
			&obj.Id,
			&obj.FirstName,
			&obj.LastName,
			&obj.Password,
			&obj.PhoneNumber,
			&obj.ImageUrl,
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
	// So'rovni qurishni boshlash
	query := `UPDATE "users" SET `
	params := []interface{}{}
	counter := 1

	// So'rovga maydonlarni qo'shish
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

	if req.Password != nil {
		query += `password_hash = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.Password)
		counter++
		updated = true
	}

	if req.ImageUrl != nil {
		query += `image_url = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.ImageUrl)
		counter++
		updated = true
	}

	if req.PhoneNumber != nil {
		query += `phone_number = $` + fmt.Sprint(counter) + `, `
		params = append(params, *req.PhoneNumber)
		counter++
		updated = true
	}

	// Agar hech qanday maydon yangilanmagan bo'lsa, xato qaytarish
	if !updated {
		return 0, fmt.Errorf("yangilanish uchun hech qanday maydon kiritilmagan")
	}

	// Oxirgi vergulani olib tashlash va `updated_at` maydonini qo'shish
	query = query[:len(query)-2] + `, updated_at = now()`

	// WHERE shartini qo'shish
	query += ` WHERE id = $` + fmt.Sprint(counter)
	params = append(params, req.Id)

	// So'rovni bajarish
	result, err := u.db.Exec(ctx, query, params...)
	if err != nil {
		return 0, err
	}

	// Qatorlar sonini olish
	rowsAffected := result.RowsAffected()

	// Qaytarish
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

func (u *userRepo) GetByPhone(ctx context.Context, login *models.Login) (*models.User, error) {
	res := &models.User{}
	query := `
        SELECT
            id,
            first_name,
            last_name,
            password_hash,
            phone_number,
            created_at,
            updated_at
        FROM
            "users"
        WHERE
            phone_number = $1`

	err := u.db.QueryRow(ctx, query, login.PhoneNumber).Scan(
		&res.Id,
		&res.FirstName,
		&res.LastName,
		&res.Password,
		&res.PhoneNumber,
		&res.CreatedAt, // created_at as time.Time
		&res.UpdatedAt, // updated_at as time.Time
	)
	if err != nil {
		return res, err
	}

	return res, nil
}
