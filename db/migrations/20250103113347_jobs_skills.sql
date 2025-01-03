-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS jobs_skills (
    job_id int,
    skill_id int,
    PRIMARY KEY (job_id, skill_id),
    FOREIGN KEY (job_id) REFERENCES jobs(id) ON DELETE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skills(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE jobs_skills;
-- +goose StatementEnd
