package models

type PrimaryKey struct {
	Id int `json:"id"`
}

type Role struct {
	Id          string `json:"id"`
	RoleName    string `json:"name"`
	Description string `json:"description"`
}

type CreateRole struct {
	RoleName    string `json:"name"`
	Description string `json:"description"`
}

type UpdateRole struct {
	Id          string `json:"id"`
	RoleName    string `json:"name"`
	Description string `json:"description"`
}

type GetListRoleRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListRoleResponse struct {
	Count int     `json:"count"`
	Roles []*Role `json:"roles"`
}
