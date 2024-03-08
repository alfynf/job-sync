-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    uuid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar NOT NULL,
    birthdate varchar(10),
    gender integer,
    phone varchar,
    profile_picture varchar,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
