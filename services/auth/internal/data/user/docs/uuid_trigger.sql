CREATE TRIGGER IF NOT EXISTS UUIDv4_generate
AFTER INSERT ON Users
FOR EACH ROW
WHEN (NEW.id IS NULL)
BEGIN
UPDATE Users SET id = (select lower(hex( randomblob(4)) || '-' ||      hex( randomblob(2))
    || '-' || '4' || substr( hex( randomblob(2)), 2) || '-'
    || substr('AB89', 1 + (abs(random()) % 4) , 1)  ||
    substr(hex(randomblob(2)), 2) || '-' || hex(randomblob(6))) ) WHERE rowid = NEW.rowid;
END;