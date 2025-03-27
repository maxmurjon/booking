package postgres

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRoleRepo struct {
	db *pgxpool.Pool
}

func (u *userRoleRepo) Create(ctx context.Context, req *models.UserRole) (*models.PrimaryKeyUUID, error) {
	query := `
		INSERT INTO user_roles (
			user_id,
			role_id
		) VALUES ($1,$2)
		RETURNING user_id;
	`

	var newID string
	err := u.db.QueryRow(ctx, query, req.UserID, req.RoleID).Scan(&newID)
	if err != nil {
		return nil, err
	}

	pKey := &models.PrimaryKeyUUID{
		ID: newID,
	}

	return pKey, nil
}

// GetByID retrieves a attribute by its ID.
func (u *userRoleRepo) GetByID(ctx context.Context, req *models.PrimaryKeyUUID) (*models.UserRole, error) {
	res := &models.UserRole{}
	query := `SELECT
		user_id,
		role_id
	FROM
		user_roles
	WHERE
		user_id = $1`

	err := u.db.QueryRow(ctx, query, req.ID).Scan(
		&res.UserID,
		&res.RoleID,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of attributes with pagination and optional search functionality.
func (u *userRoleRepo) GetList(ctx context.Context, req *models.GetListUserRoleRequest) (*models.GetListUserRoleResponse, error) {
	res := &models.GetListUserRoleResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		user_id,
		role_id
	FROM
		user_roles`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on attribute_name only
	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND name ILIKE '%' || :search || '%'"
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
	cQ := `SELECT count(1) FROM user_roles` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving attributes
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.UserRole{}
		err = rows.Scan(
			&obj.UserID,
			&obj.RoleID,
		)
		if err != nil {
			return res, err
		}

		res.UserRoles = append(res.UserRoles, obj)
	}

	return res, nil
}

// Update updates a attribute in the attributes table.
func (u *userRoleRepo) Update(ctx context.Context, req *models.UserRole) (int64, error) {
	query := `UPDATE user_roles SET
		user_id = :user_id,
		role_id = :role_id
	WHERE
		user_id = :user_id`

	params := map[string]interface{}{
		"user_id": req.UserID,
		"role_id": req.RoleID,
	}

	q, arr := helper.ReplaceQueryParams(query, params)
	result, err := u.db.Exec(ctx, q, arr...)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}

// Delete deletes a attribute from the attributes table by its ID.
func (u *userRoleRepo) Delete(ctx context.Context, req *models.PrimaryKeyUUID) (int64, error) {
	query := `DELETE FROM user_roles WHERE user_id = $1`

	result, err := u.db.Exec(ctx, query, req.ID)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
