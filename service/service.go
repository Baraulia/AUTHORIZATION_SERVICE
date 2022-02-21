package service

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type Authorization interface {
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type RoleList interface {
	GetRoleById(id int) (*model.Roles, error)
	SelectPermissionByRoleId(id int) ([]model.Permission, error)
	CreateRole(role *model.Role) (int, error)
	CreatePermission(permission *model.Permission) (int, error)
	CreateRoleToPermission(rp *model.RoleToPermission) error
	BindUserWithRole(user *auth_proto.User) error
}

type Service struct {
	Authorization
	RoleList
}

func NewService(rep *repository.Repository, logger logging.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(*rep, logger),
		RoleList:      NewRoleListService(*rep, logger),
	}
}
