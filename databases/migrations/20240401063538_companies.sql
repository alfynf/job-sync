-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS companies (
    uuid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(100) NOT NULL,
    establish_at varchar(4) NOT NULL,
    location varchar(50) NOT NULL,
    logo varchar,
    description text NOT NULL,
    address varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    phone varchar,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS companies;
-- +goose StatementEnd
