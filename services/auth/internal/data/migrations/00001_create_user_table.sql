-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "User" (
    id VARCHAR(36) PRIMARY KEY,
    full_name VARCHAR(255),
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash TEXT,
    created_at DATETIME DEFAULT (CURRENT_TIMESTAMP),
    updated_at DATETIME DEFAULT (CURRENT_TIMESTAMP)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "User";
-- +goose StatementEnd
