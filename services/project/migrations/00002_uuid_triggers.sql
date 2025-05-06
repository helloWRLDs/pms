-- +goose Up
-- +goose StatementBegin
CREATE TRIGGER IF NOT EXISTS UUIDv4_generate_Project
AFTER INSERT ON Project
FOR EACH ROW
WHEN (NEW.id IS NULL)
BEGIN
    UPDATE Project SET id = (
        SELECT LOWER(HEX(RANDOMBLOB(4)) || '-' || HEX(RANDOMBLOB(2)) || '-' ||
                     '4' || SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                     SUBSTR('AB89', 1 + (ABS(RANDOM()) % 4), 1) ||
                     SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                     HEX(RANDOMBLOB(6)))
    ) WHERE rowid = NEW.rowid;
END;

CREATE TRIGGER IF NOT EXISTS UUIDv4_generate_Task
AFTER INSERT ON Task
FOR EACH ROW
WHEN (NEW.id IS NULL)
BEGIN
    UPDATE Task SET id = (
        SELECT LOWER(HEX(RANDOMBLOB(4)) || '-' || HEX(RANDOMBLOB(2)) || '-' ||
                     '4' || SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                     SUBSTR('AB89', 1 + (ABS(RANDOM()) % 4), 1) ||
                     SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                     HEX(RANDOMBLOB(6)))
    ) WHERE rowid = NEW.rowid;
END;

CREATE TRIGGER IF NOT EXISTS UUIDv4_generate_SubTask
AFTER INSERT ON SubTask
FOR EACH ROW
WHEN (NEW.parent_id IS NULL OR NEW.child_id IS NULL)
BEGIN
    UPDATE SubTask SET 
        parent_id = COALESCE(NEW.parent_id, (
            SELECT LOWER(HEX(RANDOMBLOB(4)) || '-' || HEX(RANDOMBLOB(2)) || '-' ||
                         '4' || SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                         SUBSTR('AB89', 1 + (ABS(RANDOM()) % 4), 1) ||
                         SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                         HEX(RANDOMBLOB(6)))
        )),
        child_id = COALESCE(NEW.child_id, (
            SELECT LOWER(HEX(RANDOMBLOB(4)) || '-' || HEX(RANDOMBLOB(2)) || '-' ||
                         '4' || SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                         SUBSTR('AB89', 1 + (ABS(RANDOM()) % 4), 1) ||
                         SUBSTR(HEX(RANDOMBLOB(2)), 2) || '-' ||
                         HEX(RANDOMBLOB(6)))
        ))
    WHERE rowid = NEW.rowid;
END;

CREATE TRIGGER IF NOT EXISTS UUIDv4_generate_Sprint
AFTER INSERT ON Sprint
FOR EACH ROW
WHEN (NEW.id IS NULL)
BEGIN
    UPDATE Sprint SET id = (
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
DROP TRIGGER IF EXISTS UUIDv4_generate_Project;
DROP TRIGGER IF EXISTS UUIDv4_generate_Task;
DROP TRIGGER IF EXISTS UUIDv4_generate_SubTask;
DROP TRIGGER IF EXISTS UUIDv4_generate_Sprint;
-- +goose StatementEnd
