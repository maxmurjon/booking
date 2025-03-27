package models

type RolePermission struct {
	RoleID int `json:"role_id"`
	PermissionId    int `json:"permission_id"`
}

type GetListRolePermissionRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListRolePermissionResponse struct {
	Count int     `json:"count"`
	RolePermissions []*RolePermission `json:"role_permission"`
}
