package database

import (
	"database/sql"
	"fmt"
	"github.com/Baraulia/AUTHENTICATION_SERVICE/pkg/logging"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
	logger   logging.Logger
}

func NewPostgresDB(database PostgresDB) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		database.Username, database.Password, database.Host, database.Port, database.DBName, database.SSLMode))
	if err != nil {

		return nil, fmt.Errorf("error connecting to database:%s", err)
	}
	err = db.Ping()
	if err != nil {
		database.logger.Errorf("DB ping error:%s", err)
		return nil, err
	}
	_, err = db.Exec(ROLE_SCHEMA)
	if err != nil {
		database.logger.Errorf("Error executing initial migration into roles:%s", err)
		return nil, fmt.Errorf("error executing initial migration into roles:%s", err)
	}
	_, err = db.Exec(PERMISSION_SCHEMA)
	if err != nil {
		database.logger.Errorf("Error executing initial migration into permission:%s", err)
		return nil, fmt.Errorf("error executing initial migration into permission:%s", err)
	}
	_, err = db.Exec(REFERENCE_SCHEMA)
	if err != nil {
		database.logger.Errorf("Error executing initial migration into permission:%s", err)
		return nil, fmt.Errorf("error executing initial migration into permission:%s", err)
	}
	return db, nil
}

const ROLE_SCHEMA = `
	CREATE TABLE IF NOT EXISTS roles (
		id serial not null primary key,
		name varchar(225) NOT NULL UNIQUE
	);
`

const PERMISSION_SCHEMA = `
	CREATE TABLE IF NOT EXISTS permissions (
		id serial not null primary key,
		description varchar(225) NOT NULL UNIQUE
	);
`
const REFERENCE_SCHEMA = `
		CREATE TABLE IF NOT EXISTS role_permissions (
		role_id int references roles(id) on delete cascade,
		permission_id int references permissions(id) on delete cascade,
		PRIMARY KEY(role_id, permission_id)
	);
	CREATE TABLE IF NOT EXISTS usersss (
		id serial not null,
		role_id int not null,
		user_id int not null,
		PRIMARY KEY (id),
		FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);
`