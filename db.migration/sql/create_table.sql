
	CREATE TABLE IF NOT EXISTS roles (
		id serial not null primary key,
		name varchar(225) NOT NULL UNIQUE
	);


	CREATE TABLE IF NOT EXISTS permissions (
		id serial not null primary key,
		description varchar(225) NOT NULL UNIQUE
	);


	CREATE TABLE IF NOT EXISTS role_permissions (
		role_id int references roles(id) on delete cascade,
		permission_id int references permissions(id) on delete cascade,
		CONSTRAINT role_permissions_pkey PRIMARY KEY(role_id, permission_id)
	);


	CREATE TABLE IF NOT EXISTS user_role (
        role_id int DEFAULT 1 references roles(id) on delete SET DEFAULT,
        user_id int,
        CONSTRAINT role_user_pkey PRIMARY KEY(role_id, user_id)
);