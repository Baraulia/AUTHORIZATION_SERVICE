package repository

import (
	"database/sql"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/model"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go

type Authorization interface {
	GetUser(email, password string) (*model.User, error)
}

type RoleList interface {
	GetById(id int) (*model.Roles, error)
	SelectPermission (id int) []model.Permission
	CreateRole(role *model.Role) (*model.Role, error)
	CreatePermission(permission *model.Permission, role int)(*model.Permission, error)
	CreateRoleToPermission(rp *model.RoleToPermission)(*model.RoleToPermission, error)
}

type Repository struct {
	Authorization
	RoleList
}

func NewRepository(db *sql.DB, logger logging.Logger) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db, logger),
		RoleList:      NewRoleListPostgres(db, logger),
	}
}
