package service

import (
	authProto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/service_mock.go

type Authorization interface {
	GenerateTokensByAuthUser(user *authProto.User) (*authProto.GeneratedTokens, error)
	ParseToken(token string) (*authProto.UserRole, error)
	RefreshTokens(refreshToken string) (*authProto.GeneratedTokens, error)
	CheckRoleRights(neededPerms []string, neededRole string, givenPerms string, givenRole string) error
}

type RolePerm interface {
	GetRoleById(id int) (*model.Role, error)
	GetAllRoles() ([]model.Role, error)
	CreateRole(role string) (int, error)
	BindRoleWithPerms(rp *model.BindRoleWithPermission) error
	GetRoleByUserId(userId int) (*model.Role, error)

	GetPermsByRoleId(id int) ([]model.Permission, error)
	CreatePermission(permission string) (int, error)
	GetAllPerms() ([]model.Permission, error)

	AddRoleToUser(user *authProto.User) (bool, error)
}

type Service struct {
	RolePerm
	Authorization
}

func NewService(rep *repository.Repository, logger logging.Logger) *Service {
	return &Service{
		RolePerm:      NewRolePermService(*rep, logger),
		Authorization: NewAuthService(*rep, logger),
	}
}
