package model

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateRole struct {
	Name string `json:"name"`
}

type CreatePerm struct {
	Name string `json:"name"`
}

type Permission struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ResponseRoles struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Permissions []Permission `json:"permissions"`
}

type BindRoleWithPermission struct {
	RoleId        int   `json:"role_id"`
	PermissionsId []int `json:"permissions_id"`
}

type ResponseUser struct {
	Id          int
	Roles       []string
	Permissions []string
}

type ListRoles struct {
	Roles []Role
}
type ListPerms struct {
	Perms []Permission
}

type ErrorResponse struct {
	Message string `json:"message"`
}

type UserRole struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	Perms  string `json:"perms"`
}

type Claims struct {
	UserId    int
	RoleId    int
	ExpiresAt int64
}
