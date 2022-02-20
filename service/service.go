package service

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type Authorization interface {
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type RoleList interface {
	GetById(id int) (*model.Roles, error)
	SelectPermission(id int) ([]model.Permission, error)
	CreateRole(role *model.Role) (*model.Role, error)
	CreatePermission(permission *model.Permission) (*model.Permission, error)
	CreateRoleToPermission(rp *model.RoleToPermission) (*model.RoleToPermission, error)
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
