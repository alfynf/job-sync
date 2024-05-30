-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS job_vacancies (
    uuid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    title varchar(255) NOT NULL,
    location varchar(50) NOT NULL,
    requirement TEXT NOT NULL,
    job_type int NOT NULL,
    work_model int NOT NULL,
    end_date varchar NOT NULL,
    status int NOT NULL,
    company_uuid uuid,
    employee_uuid uuid,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_job_vacancy_company
        FOREIGN KEY (company_uuid)
        REFERENCES companies(uuid)
        ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT fk_job_vacancy_created_by
        FOREIGN KEY (employee_uuid)
        REFERENCES employees(uuid)
        ON DELETE SET NULL ON UPDATE SET NULL    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS job_vacancies;
-- +goose StatementEnd

