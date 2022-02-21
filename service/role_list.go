package service

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/repository"
)

type RoleListService struct {
	logger logging.Logger
	repo   repository.Repository
}

func NewRoleListService(repo repository.Repository, logger logging.Logger) *RoleListService {
	return &RoleListService{repo: repo, logger: logger}
}

func (s *RoleListService) GetRoleById(id int) (*model.Roles, error) {
	g, err := s.repo.RoleList.GetRoleById(id)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (s *RoleListService) SelectPermissionByRoleId(id int) ([]model.Permission, error) {
	return s.repo.RoleList.SelectPermissionByRoleId(id)
}

func (s *RoleListService) CreateRole(role *model.Role) (int, error) {
	result, err := s.repo.RoleList.CreateRole(role)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *RoleListService) CreatePermission(permission *model.Permission) (int, error) {
	result, err := s.repo.RoleList.CreatePermission(permission)
	if err != nil {
		return 0, err
	}
	return result, nil
}

func (s *RoleListService) CreateRoleToPermission(permission *model.RoleToPermission) error {
	err := s.repo.RoleList.CreateRoleToPermission(permission)
	if err != nil {
		return err
	}
	return nil
}

func (s *RoleListService) BindUserWithRole(user *auth_proto.User) error {
	return s.repo.BindUserWithRole(user)
}
