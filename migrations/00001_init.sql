-- +goose Up
CREATE TABLE IF NOT EXISTS todo(
		id serial PRIMARY KEY,
		title VARCHAR (50) UNIQUE NOT NULL,
		text TEXT NOT NULL 
	);

-- +goose Down
DROP TABLE todo;