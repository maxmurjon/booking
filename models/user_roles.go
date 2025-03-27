package models

type PrimaryKeyUUID struct{
	ID string `json:"id"`
}
type UserRole struct {
	UserID    string `json:"user_id"`
	RoleID string `json:"role_id"`
}

type GetListUserRoleRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListUserRoleResponse struct {
	Count int     `json:"count"`
	UserRoles []*UserRole `json:"user_roles"`
}
