package service

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/model"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/repository"
)

type RoleListService struct {
	logger logging.Logger
	repo repository.Repository
}

func NewRoleListService(repo repository.Repository, logger logging.Logger) *RoleListService {
	return &RoleListService{repo: repo, logger: logger}
}

func (s *RoleListService) GetById(id int) (*model.Roles, error) {
	g, err := s.repo.RoleList.GetById(id)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (s *RoleListService) SelectPermission(id int) ([]model.Permission,error) {
	return s.repo.RoleList.SelectPermission(id)
}

func (s *RoleListService) CreateRole(role *model.Role) (*model.Role, error) {
	result, err := s.repo.RoleList.CreateRole(role)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *RoleListService) CreatePermission(permission *model.Permission) (*model.Permission, error) {
	result, err := s.repo.RoleList.CreatePermission(permission)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *RoleListService) CreateRoleToPermission(permission *model.RoleToPermission) (*model.RoleToPermission, error) {
	result, err := s.repo.RoleList.CreateRoleToPermission(permission)
	if err != nil {
		return nil, err
	}
	return result, nil
}