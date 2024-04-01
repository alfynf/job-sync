-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS company_positions (
    uuid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name varchar(100) NOT NULL,
    company_uuid uuid,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_position_company
        FOREIGN KEY (company_uuid)
        REFERENCES companies(uuid)
        ON DELETE CASCADE ON UPDATE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS company_positions;
-- +goose StatementEnd
