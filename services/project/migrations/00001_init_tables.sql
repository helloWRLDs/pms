-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "Project" (
    "id" VARCHAR PRIMARY KEY,
    "title" TEXT,
    "status" VARCHAR,
    "description" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "company_id" VARCHAR,
    "progress" INT
);

CREATE TABLE IF NOT EXISTS "Task" (
    "id" VARCHAR PRIMARY KEY,
    "title" TEXT,
    "body" TEXT,
    "project_id" VARCHAR DEFAULT NULL,
    "sprint_id" VARCHAR DEFAULT NULL,
    "status" TEXT,
    "priority" INTEGER,
    "due_date" DATE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("project_id") REFERENCES "Project"("id"),
    FOREIGN KEY ("sprint_id") REFERENCES "Sprint"("id")
);

CREATE TABLE IF NOT EXISTS "SubTask" (
    "parent_id" VARCHAR,
    "child_id" VARCHAR,
    PRIMARY KEY("parent_id", "child_id"),
    FOREIGN KEY ("parent_id") REFERENCES "Task"("id"),
    FOREIGN KEY ("child_id") REFERENCES "Task"("id")
);

CREATE TABLE IF NOT EXISTS "Sprint" (
    "id" VARCHAR PRIMARY KEY,
    "title" TEXT NOT NULL,
    "description" TEXT,
    "start_date" DATE,
    "end_date" DATE,
    "project_id" VARCHAR,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("project_id") REFERENCES "Project"("id")
);

CREATE TABLE IF NOT EXISTS "TaskAssignment" (
    "user_id" VARCHAR,
    "task_id" VARCHAR,
    PRIMARY KEY("user_id", "task_id"),
    FOREIGN KEY ("task_id") REFERENCES "Task"("id")
);

CREATE TABLE IF NOT EXISTS "TaskComment" (
    "id" VARCHAR PRIMARY KEY,
    "task_id" VARCHAR NOT NULL,
    "user_id" VARCHAR,
    "body" TEXT,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("task_id") REFERENCES "Task"("id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS TaskComment;
DROP TABLE IF EXISTS TaskAssignment;
DROP TABLE IF EXISTS Sprint;
DROP TABLE IF EXISTS SubTask;
DROP TABLE IF EXISTS Task;
DROP TABLE IF EXISTS Project;
-- +goose StatementEnd