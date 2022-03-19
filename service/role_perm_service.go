package service

import (
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/repository"
)

type AuthUserService struct {
	logger logging.Logger
	repo   repository.Repository
}

func NewAuthUserService(repo repository.Repository, logger logging.Logger) *AuthUserService {
	return &AuthUserService{repo: repo, logger: logger}
}

func (a *AuthUserService) GetRoleById(id int) (*model.Role, error) {
	return a.repo.RolePerm.GetRoleById(id)
}

func (a *AuthUserService) GetRoleByUserId(userId int) (*model.Role, error) {
	return a.repo.RolePerm.GetRoleByUserId(userId)
}

func (a *AuthUserService) GetAllRoles() ([]model.Role, error) {
	return a.repo.GetAllRoles()
}

func (a *AuthUserService) CreateRole(role string) (int, error) {
	return a.repo.RolePerm.CreateRole(role)
}

func (a *AuthUserService) BindRoleWithPerms(rp *model.BindRoleWithPermission) error {
	err := a.repo.RolePerm.BindRoleWithPerms(rp)
	if err != nil {
		return err
	}
	return nil
}

func (a *AuthUserService) GetPermsByRoleId(id int) ([]model.Permission, error) {
	return a.repo.RolePerm.GetPermsByRoleId(id)
}

func (a *AuthUserService) CreatePermission(permission string) (int, error) {
	return a.repo.RolePerm.CreatePermission(permission)
}

func (a *AuthUserService) GetAllPerms() ([]model.Permission, error) {
	return a.repo.RolePerm.GetAllPerms()
}

func (a *AuthUserService) AddRoleToUser(user *authProto.User) (bool, error) {
	err := a.repo.AddRoleToUser(user)
	if err != nil {
		return false, err
	}
	return true, nil
}
