-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Role" (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT UNIQUE NOT NULL,
	permissions TEXT NOT NULL DEFAULT '[]' -- Store JSON as TEXT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "Role";
-- +goose StatementEnd
