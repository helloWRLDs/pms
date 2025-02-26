-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER IF NOT EXISTS UUIDv4_generate_Users
AFTER INSERT ON User
FOR EACH ROW
WHEN (NEW.id IS NULL)
BEGIN
    UPDATE User SET id = (
        SELECT LOWER(HEX(RANDOMBLOB(4)) || '-' || HEX(RANDOMBLOB(2)) || '-' ||
                     '4' || SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                     SUBSTR('AB89', 1 + (ABS(RANDOM()) % 4), 1) ||
                     SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                     HEX(RANDOMBLOB(6)))
    ) WHERE rowid = NEW.rowid;
END;
CREATE TRIGGER IF NOT EXISTS UUIDv4_generate_Company
AFTER INSERT ON Company
FOR EACH ROW
WHEN (NEW.id IS NULL)
BEGIN
    UPDATE Company SET id = (
        SELECT LOWER(HEX(RANDOMBLOB(4)) || '-' || HEX(RANDOMBLOB(2)) || '-' ||
                     '4' || SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                     SUBSTR('AB89', 1 + (ABS(RANDOM()) % 4), 1) ||
                     SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                     HEX(RANDOMBLOB(6)))
    ) WHERE rowid = NEW.rowid;
END;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TRIGGER UUIDv4_generate_Users;
DROP TRIGGER UUIDv4_generate_Company;
-- +goose StatementEnd
