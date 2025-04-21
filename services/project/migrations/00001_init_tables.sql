-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Project" (
	"id" VARCHAR UNIQUE,
	"title" VARCHAR,
	"description" TEXT,
	"status" VARCHAR(50),
	"phone" VARCHAR,
	"company_id" VARCHAR,
	"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "Project";
-- +goose StatementEnd
