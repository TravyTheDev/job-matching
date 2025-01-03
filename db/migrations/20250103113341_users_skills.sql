-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users_skills (
    user_id int,
    skill_id int,
    PRIMARY KEY (user_id, skill_id),
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY (skill_id) REFERENCES skills(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users_skills;
-- +goose StatementEnd
