package service

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/model"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type Authorization interface {
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type RoleList interface {
	GetById(id int) (*model.Roles, error)
	SelectPermission (id int) []model.Permission
	CreateRole(role *model.Role) (*model.Role, error)
	CreatePermission(permission *model.Permission)(*model.Permission, error)
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
