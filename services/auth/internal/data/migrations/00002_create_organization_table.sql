-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Organization" (
	id VARCHAR(36) PRIMARY KEY UNIQUE,
	name TEXT NOT NULL,
	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
	updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "Organization";
-- +goose StatementEnd
