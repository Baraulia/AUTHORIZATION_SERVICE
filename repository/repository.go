package repository

import (
	"database/sql"
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	_ "stlab.itechart-group.com/go/food_delivery/authorization_service/docs"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
)

//go:generate mockgen -source=repository.go -destination=mocks/repository_mock.go

type RolePerm interface {
	GetRoleById(id int) (*model.Role, error)
	GetAllRoles() ([]model.Role, error)
	CreateRole(role string) (int, error)
	BindRoleWithPerms(rp *model.BindRoleWithPermission) error
	GetRoleByUserId(userId int) (*model.Role, error)
	GetRoleByName(roleName string) (*model.Role, error)

	GetPermsByRoleId(id int) ([]model.Permission, error)
	CreatePermission(permission string) (int, error)
	GetAllPerms() ([]model.Permission, error)

	AddRoleToUser(user *authProto.User) error
}

type Repository struct {
	RolePerm
}

func NewRepository(db *sql.DB, logger logging.Logger) *Repository {
	return &Repository{
		RolePerm: NewRolePermPostgres(db, logger),
	}
}
