-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS employees (
    uuid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name varchar(50) NOT NULL,
    last_name varchar(50) NOT NULL,
    username varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar NOT NULL,
    profile_picture varchar,
    company_uuid uuid,
    position_uuid uuid,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_employee_company
        FOREIGN KEY (company_uuid)
        REFERENCES companies(uuid)
        ON DELETE SET NULL ON UPDATE SET NULL,
    CONSTRAINT fk_employee_position
        FOREIGN KEY (position_uuid)
        REFERENCES company_positions(uuid)
        ON DELETE SET NULL ON UPDATE SET NULL    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS employees;
-- +goose StatementEnd

