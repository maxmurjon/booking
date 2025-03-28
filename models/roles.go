package models

type PrimaryKey struct {
	Id string `json:"id"`
}

type Role struct {
	Id          string `json:"id"`
	Name    string `json:"name"`
}

type CreateRole struct {
	Name    string `json:"name"`
}

type UpdateRole struct {
	Id          string `json:"id"`
	Name    string `json:"name"`
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
