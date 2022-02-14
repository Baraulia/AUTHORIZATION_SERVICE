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

func (r *RoleListPostgres) GetById(id int) (*model.Roles, error) {

	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("GetByID: can not starts transaction:%s", err)
		return nil, fmt.Errorf("getByID: can not starts transaction:%w", err)
	}
	var role model.Roles
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

func (r *RoleListPostgres) CreateRole(role *model.Role) (*model.Role, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("CreateRole: can not starts transaction:%s", err)
		return nil, fmt.Errorf("createRole: can not starts transaction:%w", err)
	}
	var createdRole model.Role
	defer transaction.Rollback()
	row := transaction.QueryRow("INSERT INTO roles (name) VALUES ($1) RETURNING id, name", role.Name)
	if err := row.Scan(&createdRole.ID, &createdRole.Name); err != nil {
		r.logger.Errorf("CreateRole: error while scanning for role:%s", err)
		return nil, fmt.Errorf("createRole: error while scanning for role:%w", err)
	}
	return &createdRole, transaction.Commit()
}

func (r *RoleListPostgres) CreatePermission(permission *model.Permission) (*model.Permission, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("CreatePermission: can not starts transaction:%s", err)
		return nil, fmt.Errorf("createPermission: can not starts transaction:%w", err)
	}
	var createdPerm model.Permission
	defer transaction.Rollback()
	row := transaction.QueryRow("INSERT INTO permissions (description) VALUES ($1) RETURNING id, description", permission.Description)
	if err := row.Scan(&createdPerm.ID, &createdPerm.Description); err != nil {
		r.logger.Errorf("CreatePermission: error while scanning for permission:%s", err)
		return nil, fmt.Errorf("createPermission: error while scanning for permission:%w", err)
	}
	return &createdPerm, transaction.Commit()
}