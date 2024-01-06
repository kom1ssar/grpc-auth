-- +goose Up
CREATE TABLE users(
    id SERIAL PRIMARY KEY ,
    name TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    role INTEGER DEFAULT 0,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


-- +goose Down
DROP TABLE users;
