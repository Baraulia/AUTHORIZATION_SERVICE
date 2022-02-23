package service

import (
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
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
	return s.repo.RolePerm.GetRoleById(id)
}

func (s *RolePermService) GetAllRoles() ([]model.Role, error) {
	return s.repo.GetAllRoles()
}

func (s *RolePermService) CreateRole(role string) (int, error) {
	return s.repo.RolePerm.CreateRole(role)
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
	return s.repo.RolePerm.CreatePermission(permission)
}

func (s *RolePermService) GetAllPerms() ([]model.Permission, error) {
	return s.repo.RolePerm.GetAllPerms()
}

func (s *RolePermService) AddRoleToUser(user *authProto.User) error {
	return s.repo.AddRoleToUser(user)
}
