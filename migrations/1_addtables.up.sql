CREATE EXTENSION IF NOT EXISTS "citext";

CREATE TABLE users (
	id SERIAL NOT NULL,
	email CITEXT UNIQUE NOT NULL CHECK (char_length(email) <= 254),
	password BYTEA NOT NULL, -- Password Bcrypt
	createdat DATE NOT NULL DEFAULT NOW(),
	updatedat DATE NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

CREATE TABLE song (
	id SERIAL NOT NULL,
	name varchar(255) NOT NULL,
	ownerid INT REFERENCES users(id) ON DELETE CASCADE,
	uploadat DATE NOT NULL DEFAULT NOW(),
	filehash varchar(64) NOT NULL,
	filepath TEXT NOT NULL,
	size BIGINT NOT NULL,
	PRIMARY KEY (id)
);