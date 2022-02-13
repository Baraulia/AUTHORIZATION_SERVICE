package service

import (
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/model"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/repository"
)

type RoleListService struct {
	logger logging.Logger
	repo repository.RoleList
}

func NewRoleListService(repo repository.RoleList, logger logging.Logger) *RoleListService {
	return &RoleListService{repo: repo, logger: logger}
}

func (s *RoleListService) GetById(id int) (*model.Role, error) {
	return s.repo.GetById(id)
}

func (s *RoleListService) SelectPermission(id int) []model.Permission {
	return s.repo.SelectPermission(id)
}
