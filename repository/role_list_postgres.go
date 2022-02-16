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
		return nil, fmt.Errorf("GetByID: can not starts transaction:%w", err)
	}
	var role model.Roles
	query := "SELECT id, name FROM roles WHERE id = $1"
	row := transaction.QueryRow(query, id)
	if err := row.Scan(&role.ID, &role.Name); err != nil {
		logrus.Errorf("GetByID: error while scanning for book:%s", err)
		return nil, fmt.Errorf("GetByID: repository error:%w", err)
	}
	role.Permissions, err = r.SelectPermission(role.ID)
	if err != nil {
		return nil, fmt.Errorf("error while getting bound roles:%w", err)
	}
	return &role, transaction.Commit()
}


func (r *RoleListPostgres) SelectPermission(id int) ([]model.Permission, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("ReturnPermission: can not starts transaction:%s", err)
		return nil, fmt.Errorf("ReturnPermission: can not starts transaction:%s", err)
	}
	var permissions []model.Permission
	query := "SELECT id, description FROM permissions JOIN role_permissions ON permissions.id = role_permissions.permission_id AND role_permissions.role_id = $1"
	rows, err := transaction.Query(query, id)
	if err != nil {
		logrus.Errorf("ReturnPermission: can not executes a query:%s", err)
		return nil, fmt.Errorf("ReturnPermission: repository error:%w", err)
	}
	for rows.Next() {
		var permission model.Permission
		if err := rows.Scan(&permission.ID, &permission.Description); err != nil {
			logrus.Errorf("ReturnPermission: error while scanning :%s", err)
			return nil, fmt.Errorf("ReturnPermission: repository error:%w", err)
		}
		permissions = append(permissions, permission)
	}
	return permissions, transaction.Commit()
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

func (r *RoleListPostgres) CreateRoleToPermission(rp *model.RoleToPermission) (*model.RoleToPermission, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("CreateRP: can not starts transaction:%s", err)
		return nil, fmt.Errorf("createRP: can not starts transaction:%w", err)
	}
	var createdRP model.RoleToPermission
	defer transaction.Rollback()
	row := transaction.QueryRow("INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2) RETURNING role_id, permission_id", rp.RoleId, rp.PermissionId )
	if err := row.Scan(&createdRP.RoleId, &createdRP.PermissionId); err != nil {
		r.logger.Errorf("CreateRP: error while scanning for permission:%s", err)
		return nil, fmt.Errorf("createRP: error while scanning for permission:%w", err)
	}
	return &createdRP, transaction.Commit()
}

