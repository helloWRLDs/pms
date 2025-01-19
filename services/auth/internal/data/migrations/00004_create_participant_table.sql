-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Participant" (
	org_id VARCHAR(36) NOT NULL,
	user_id VARCHAR(36) NOT NULL,
	role_id INTEGER NOT NULL,
	PRIMARY KEY (org_id, user_id, role_id),
	FOREIGN KEY (org_id) REFERENCES Organization(id) ON DELETE CASCADE,
	FOREIGN KEY (user_id) REFERENCES User(id) ON DELETE CASCADE,
	FOREIGN KEY (role_id) REFERENCES Role(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "Participant";
-- +goose StatementEnd
