package repository

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	auth_proto "stlab.itechart-group.com/go/food_delivery/authorization_service/GRPC"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/model"
	"stlab.itechart-group.com/go/food_delivery/authorization_service/pkg/logging"
)

type RolePermPostgres struct {
	logger logging.Logger
	db     *sql.DB
}

func NewRolePermPostgres(db *sql.DB, logger logging.Logger) *RolePermPostgres {
	return &RolePermPostgres{db: db, logger: logger}
}

func (r *RolePermPostgres) GetRoleById(id int) (*model.Role, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("GetRoleById: can not starts transaction:%s", err)
		return nil, fmt.Errorf("GetRoleById: can not starts transaction:%w", err)
	}
	var role model.Role
	query := "SELECT id, name FROM roles WHERE id = $1"
	row := transaction.QueryRow(query, id)
	if err := row.Scan(&role.ID, &role.Name); err != nil {
		logrus.Errorf("GetRoleById: error while scanning for role:%s", err)
		return nil, fmt.Errorf("GetRoleById: repository error:%w", err)
	}
	return &role, transaction.Commit()
}

func (r *RolePermPostgres) GetAllRoles() ([]model.Role, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("GetAllRoles: can not starts transaction:%s", err)
		return nil, fmt.Errorf("GetAllRoles: can not starts transaction:%w", err)
	}
	var roles []model.Role
	query := "SELECT id, name FROM roles"
	rows, err := transaction.Query(query)
	if err != nil {
		logrus.Errorf("GetAllRoles: can not executes a query:%s", err)
		return nil, fmt.Errorf("GetAllRoles: repository error:%w", err)
	}
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.ID, &role.Name); err != nil {
			logrus.Errorf("GetAllRoles: error while scanning for role:%s", err)
			return nil, fmt.Errorf("GetAllRoles: repository error:%w", err)
		}
		roles = append(roles, role)
	}
	return roles, transaction.Commit()
}

func (r *RolePermPostgres) CreateRole(role string) (int, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("CreateRole: can not starts transaction:%s", err)
		return 0, fmt.Errorf("createRole: can not starts transaction:%w", err)
	}
	var roleId int
	defer transaction.Rollback()
	row := transaction.QueryRow("INSERT INTO roles (name) VALUES ($1) RETURNING id", role)
	if err := row.Scan(&roleId); err != nil {
		r.logger.Errorf("CreateRole: error while scanning for roleId:%s", err)
		return 0, fmt.Errorf("createRole: error while scanning for roleId:%w", err)
	}
	return roleId, transaction.Commit()
}

func (r *RolePermPostgres) BindRoleWithPerms(rp *model.BindRoleWithPermission) error {
	transaction, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("BindRoleWithPerms: can not starts transaction:%s", err)
		return fmt.Errorf("BindRoleWithPerms: can not starts transaction:%w", err)
	}
	defer transaction.Rollback()
	query := "INSERT INTO role_permissions (role_id, permission_id) VALUES ($1, $2)"
	for _, perm := range rp.PermissionsId {
		_, err := transaction.Exec(query, rp.RoleId, perm)
		if err != nil {
			r.logger.Errorf("BindRoleWithPerms: error while binding role:%d and permissin:%d:%s", rp.RoleId, perm, err)
			return fmt.Errorf("BindRoleWithPerms: error while binding role:%d and permissin:%d:%w", rp.RoleId, perm, err)
		}
	}
	return transaction.Commit()
}

func (r *RolePermPostgres) GetPermsByRoleId(id int) ([]model.Permission, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("GetPermsByRoleId: can not starts transaction:%s", err)
		return nil, fmt.Errorf("GetPermsByRoleId: can not starts transaction:%s", err)
	}
	var permissions []model.Permission
	query := "SELECT id, description FROM permissions JOIN role_permissions ON permissions.id = role_permissions.permission_id AND role_permissions.role_id = $1"
	rows, err := transaction.Query(query, id)
	if err != nil {
		logrus.Errorf("GetPermsByRoleId: can not executes a query:%s", err)
		return nil, fmt.Errorf("GetPermsByRoleId: repository error:%w", err)
	}
	for rows.Next() {
		var permission model.Permission
		if err := rows.Scan(&permission.ID, &permission.Name); err != nil {
			logrus.Errorf("GetPermsByRoleId: error while scanning :%s", err)
			return nil, fmt.Errorf("GetPermsByRoleId: repository error:%w", err)
		}
		permissions = append(permissions, permission)
	}
	return permissions, transaction.Commit()
}

func (r *RolePermPostgres) CreatePermission(permission string) (int, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("CreatePermission: can not starts transaction:%s", err)
		return 0, fmt.Errorf("createPermission: can not starts transaction:%w", err)
	}
	var permId int
	defer transaction.Rollback()
	row := transaction.QueryRow("INSERT INTO permissions (description) VALUES ($1) RETURNING id", permission)
	if err := row.Scan(&permId); err != nil {
		r.logger.Errorf("CreatePermission: error while scanning for permission:%s", err)
		return 0, fmt.Errorf("createPermission: error while scanning for permission:%w", err)
	}
	return permId, transaction.Commit()
}

func (r *RolePermPostgres) GetAllPerms() ([]model.Permission, error) {
	transaction, err := r.db.Begin()
	if err != nil {
		logrus.Errorf("GetAllPerms: can not starts transaction:%s", err)
		return nil, fmt.Errorf("GetAllPerms: can not starts transaction:%s", err)
	}
	var permissions []model.Permission
	query := "SELECT id, description FROM permissions"
	rows, err := transaction.Query(query)
	if err != nil {
		logrus.Errorf("GetAllPerms: can not executes a query:%s", err)
		return nil, fmt.Errorf("GetAllPerms: repository error:%w", err)
	}
	for rows.Next() {
		var permission model.Permission
		if err := rows.Scan(&permission.ID, &permission.Name); err != nil {
			logrus.Errorf("GetAllPerms: error while scanning :%s", err)
			return nil, fmt.Errorf("GetAllPerms: repository error:%w", err)
		}
		permissions = append(permissions, permission)
	}
	return permissions, transaction.Commit()
}

func (r *RolePermPostgres) BindUserWithRole(user *auth_proto.User) error {
	transaction, err := r.db.Begin()
	if err != nil {
		r.logger.Errorf("CreateRP: can not starts transaction:%s", err)
		return fmt.Errorf("createRP: can not starts transaction:%w", err)
	}
	_, err = transaction.Exec("INSERT into user_role (role_id, user_id) values ($1, $2)", user.RoleId, user.UserId)
	if err != nil {
		r.logger.Errorf("BindUserWithRole:%s", err)
		return fmt.Errorf("BindUserWithRole:%w", err)
	}
	return transaction.Commit()
}
