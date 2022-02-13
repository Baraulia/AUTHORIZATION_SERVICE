package repository

import (
	"database/sql"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/model"
)

type Authorization interface {
	GetUser(username, password string) (*model.User, error)
}

type RoleList interface {
	GetById(id int) (*model.Role, error)
	SelectPermission (id int) []model.Permission
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
