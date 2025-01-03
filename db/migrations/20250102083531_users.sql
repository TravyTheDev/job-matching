-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    username varchar(255),
    email varchar(255) UNIQUE,
    password varchar(255),
    is_admin boolean,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
