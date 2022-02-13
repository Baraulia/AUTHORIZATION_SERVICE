package repository

import (
	"database/sql"
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/model"
	"github.com/sirupsen/logrus"
)

type RoleListPostgres struct {
	logger logging.Logger
	db *sql.DB
}

func NewRoleListPostgres(db *sql.DB, logger logging.Logger) *RoleListPostgres {
	return &RoleListPostgres{db: db, logger: logger}
}

func (r *RoleListPostgres) GetById(id int) (*model.Role, error) {

	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("GetByID: can not starts transaction:%s", err)
		return nil, fmt.Errorf("getByID: can not starts transaction:%w", err)
	}
	var role model.Role
	result := transaction.QueryRow("SELECT id, name FROM roles WHERE id = $1", id)
	if err := result.Scan(&role.ID, &role.Name); err != nil {
		logrus.Errorf("GetByID: error while scanning for user:%s", err)
		return nil, fmt.Errorf("getByID: repository error:%w", err)
	}
	role.Permissions = r.SelectPermission(role.ID)
	return &role, transaction.Commit()
}

func (r *RoleListPostgres) SelectPermission(id int) []model.Permission {
	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("GetByID: can not starts transaction:%s", err)
		return nil
	}
	var permissions []model.Permission
	var permission model.Permission
	result := transaction.QueryRow("SELECT id, description FROM permissions JOIN role_permissions ON permissions.id = role_permissions.permission_id AND role_permissions.role_id = $1", id)
	if err := result.Scan(&permission.ID, &permission.Description); err != nil {
		logrus.Errorf("GetByID: error while scanning for user:%s", err)
		return nil
	}
	permissions = append(permissions, permission)
	transaction.Commit()
	return permissions
}
