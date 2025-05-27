-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS "Document" (
    id UUID UNIQUE DEFAULT uuid_generate_v4(),
    title TEXT,
    body bytea,
    project_id UUID,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY("id")
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "Document";
-- +goose StatementEnd
