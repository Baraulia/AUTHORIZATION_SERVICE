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
	GetById(id int) (*model.Role, error)
	SelectPermission (id int) []model.Permission
}

type Service struct {
	Authorization
	RoleList
}

func NewService(repos *repository.Repository, logger logging.Logger) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization, logger),
		RoleList:      NewRoleListService(repos.RoleList, logger),
	}
}
