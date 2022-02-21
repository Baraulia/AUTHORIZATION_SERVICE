package repository

import (
	"database/sql"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go

type Authorization interface {
	GetUser(email, password string) (*model.User, error)
}

type RoleList interface {
	GetRoleById(id int) (*model.Roles, error)
	SelectPermissionByRoleId(id int) ([]model.Permission, error)
	CreateRole(role *model.Role) (int, error)
	CreatePermission(permission *model.Permission) (int, error)
	CreateRoleToPermission(rp *model.RoleToPermission) error
	BindUserWithRole(user *auth_proto.User) error
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
