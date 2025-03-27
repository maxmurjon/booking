package postgres

import (
	"comics/models"
	"comics/pkg/helper/helper"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type roleRepo struct {
	db *pgxpool.Pool
}

func (u *roleRepo) Create(ctx context.Context, req *models.CreateRole) (*models.PrimaryKey, error) {
	query := `
		INSERT INTO roles (
			name,
			description
		) VALUES ($1,$2)
		RETURNING id;
	`

	var newID int
	err := u.db.QueryRow(ctx, query, req.RoleName, req.Description).Scan(&newID)
	if err != nil {
		return nil, err
	}

	pKey := &models.PrimaryKey{
		Id: newID,
	}

	return pKey, nil
}

// GetByID retrieves a role by its ID.
func (u *roleRepo) GetByID(ctx context.Context, req *models.PrimaryKey) (*models.Role, error) {
	res := &models.Role{}
	query := `SELECT
		id,
		name,
		description
	FROM
		roles
	WHERE
		id = $1`

	err := u.db.QueryRow(ctx, query, req.Id).Scan(
		&res.Id,
		&res.RoleName,
		&res.Description,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (u *roleRepo) GetByName(ctx context.Context, req *models.PrimaryKeyUUID) (*models.Role, error) {
	res := &models.Role{}
	query := `SELECT
		id,
		name,
		description
	FROM
		roles
	WHERE
		name = $1`

	err := u.db.QueryRow(ctx, query, req.ID).Scan(
		&res.Id,
		&res.RoleName,
		&res.Description,
	)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetList retrieves a list of roles with pagination and optional search functionality.
func (u *roleRepo) GetList(ctx context.Context, req *models.GetListRoleRequest) (*models.GetListRoleResponse, error) {
	res := &models.GetListRoleResponse{}
	params := make(map[string]interface{})
	var arr []interface{}

	query := `SELECT
		id,
		name,
		description
	FROM
		roles`
	filter := " WHERE 1=1"
	offset := " OFFSET 0"
	limit := " LIMIT 10"

	// Implement search on role_name only
	if len(req.Search) > 0 {
		params["search"] = req.Search
		filter += " AND role_name ILIKE '%' || :search || '%'"
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
	cQ := `SELECT count(1) FROM roles` + filter
	cQ, arr = helper.ReplaceQueryParams(cQ, params)
	err := u.db.QueryRow(ctx, cQ, arr...).Scan(
		&res.Count,
	)
	if err != nil {
		return res, err
	}

	// Main query for retrieving roles
	q := query + filter + offset + limit
	q, arr = helper.ReplaceQueryParams(q, params)
	rows, err := u.db.Query(ctx, q, arr...)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		obj := &models.Role{}
		err = rows.Scan(
			&obj.Id,
			&obj.RoleName,
			&obj.Description,
		)
		if err != nil {
			return res, err
		}

		res.Roles = append(res.Roles, obj)
	}

	return res, nil
}

// Update updates a role in the roles table.
func (u *roleRepo) Update(ctx context.Context, req *models.UpdateRole) (int64, error) {
	query := `UPDATE roles SET
		name = $1,
		description = $2
	WHERE
		id = $3`

	// Execute the query with positional parameters
	result, err := u.db.Exec(ctx, query, req.RoleName, req.Description, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, nil
}

// Delete deletes a role from the roles table by its ID.
func (u *roleRepo) Delete(ctx context.Context, req *models.PrimaryKey) (int64, error) {
	query := `DELETE FROM roles WHERE id = $1`

	result, err := u.db.Exec(ctx, query, req.Id)
	if err != nil {
		return 0, err
	}

	rowsAffected := result.RowsAffected()

	return rowsAffected, err
}
