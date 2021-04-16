CREATE EXTENSION IF NOT EXISTS "citext";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
	id SERIAL NOT NULL,
	email CITEXT UNIQUE NOT NULL CHECK (char_length(email) <= 254),
	password BYTEA NOT NULL, -- Password Bcrypt
	createdat DATE NOT NULL DEFAULT NOW(),
	updatedat DATE NOT NULL DEFAULT NOW(),
	PRIMARY KEY (id)
);

CREATE TABLE media (
	id uuid NOT NULL DEFAULT uuid_generate_v4(),
	name varchar(255) NOT NULL,
	type TEXT NOT NULL,
	ownerid INT REFERENCES users(id) ON DELETE CASCADE,
	uploaded_time DATE NOT NULL DEFAULT NOW(),
	filehash citext NOT NULL,
	size BIGINT NOT NULL,
	PRIMARY KEY (id)
);