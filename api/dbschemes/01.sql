CREATE TABLE users (
	id SERIAL PRIMARY KEY,
	username VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	role VARCHAR(255) DEFAULT('user') NOT NULL
);
ALTER TABLE users
ADD CONSTRAINT uc_users_username UNIQUE(username);

CREATE TABLE sessions (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id),
    refresh_token UUID NOT NULL,
    ip VARCHAR(15) NOT NULL,
    expires_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT timezone('utc', now())
);
ALTER TABLE sessions
ADD CONSTRAINT uc_sessions_refreshtoken UNIQUE(refresh_token);