package service

import (
	auth_proto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/repository"
)

type RolePermService struct {
	logger logging.Logger
	repo   repository.Repository
}

func NewRolePermService(repo repository.Repository, logger logging.Logger) *RolePermService {
	return &RolePermService{repo: repo, logger: logger}
}

func (s *RolePermService) GetRoleById(id int) (*model.Role, error) {
	role, err := s.repo.RolePerm.GetRoleById(id)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (s *RolePermService) GetAllRoles() ([]model.Role, error) {
	roles, err := s.repo.RolePerm.GetAllRoles()
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (s *RolePermService) CreateRole(role string) (int, error) {
	roleId, err := s.repo.RolePerm.CreateRole(role)
	if err != nil {
		return 0, err
	}
	return roleId, nil
}

func (s *RolePermService) BindRoleWithPerms(rp *model.BindRoleWithPermission) error {
	err := s.repo.RolePerm.BindRoleWithPerms(rp)
	if err != nil {
		return err
	}
	return nil
}

func (s *RolePermService) GetPermsByRoleId(id int) ([]model.Permission, error) {
	return s.repo.RolePerm.GetPermsByRoleId(id)
}

func (s *RolePermService) CreatePermission(permission string) (int, error) {
	permId, err := s.repo.RolePerm.CreatePermission(permission)
	if err != nil {
		return 0, err
	}
	return permId, nil
}

func (s *RolePermService) GetAllPerms() ([]model.Permission, error) {
	return s.repo.RolePerm.GetAllPerms()
}

func (s *RolePermService) BindUserWithRole(user *auth_proto.User) error {
	return s.repo.BindUserWithRole(user)
}
