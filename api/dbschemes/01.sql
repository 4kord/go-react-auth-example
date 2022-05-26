CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(255) DEFAULT('user') NOT NULL
);
ALTER TABLE users ADD CONSTRAINT uc_users_username UNIQUE(username);