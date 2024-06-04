-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS applicants (
    uuid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_uuid uuid,
    job_vacancy_uuid uuid,
    notes TEXT,
    cv varchar,
    status int NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_applicant_user
        FOREIGN KEY (user_uuid)
        REFERENCES users(uuid)
        ON DELETE SET NULL ON UPDATE SET NULL,
    CONSTRAINT fk_applicant_job_vacancy
        FOREIGN KEY (job_vacancy_uuid)
        REFERENCES job_vacancies(uuid)
        ON DELETE SET NULL ON UPDATE SET NULL    
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS applicants;
-- +goose StatementEnd

