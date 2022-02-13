package repository

import (
	"database/sql"
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	"github.com/Baraulia/AUTHORIZATION_SERVICE/model"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	logger logging.Logger
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB, logger logging.Logger) *AuthPostgres {
	return &AuthPostgres{db: db, logger: logger}
}

func (r *AuthPostgres) GetUser(email, password string) (*model.User, error) {
		transaction, err := r.db.Begin()
		if err != nil {
			logrus.Errorf("GetUserByEmail: can not starts transaction:%s", err)
			return nil, fmt.Errorf("getUserByEmail: can not starts transaction:%w", err)
		}
		var User model.User
		query := "SELECT id, email, password FROM users WHERE email = $1 AND password = $2"
		row := transaction.QueryRow(query, email, password)
		if err := row.Scan(&User.ID, &User.Email, &User.Password); err != nil {
			logrus.Errorf("Error while scanning for user:%s", err)
			return nil, fmt.Errorf("getUserByEmail: repository error:%w", err)

		}
		return &User, transaction.Commit()
	}
