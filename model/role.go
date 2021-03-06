package model

type User struct {
	ID       int    `json:"id"       sql:"id"`
	Email    string `json:"email"    validate:"email"    sql:"email"`
	Password string `json:"password" validate:"password" sql:"password"`
}

type Role struct {
	ID   int    `json:"id"   sql:"id"`
	Name string `json:"name" validate:"name" sql:"name"`
}

type Permission struct {
	ID          int    `json:"id" sql:"id"`
	Description string `json:"description" validate:"description" sql:"description"`
}

type Roles struct {
	ID          int          `json:"id"   sql:"id"`
	Name        string       `json:"name" validate:"name" sql:"name"`
	Permissions []Permission `json:"permissions"`
}

type RoleToPermission struct {
	RoleId       int `json:"role_id"   sql:"role_id"`
	PermissionId int `json:"permission_id"   sql:"permission_id"`
}

type ResponseUser struct {
	Id          int
	Roles       []string
	Permissions []string
}
