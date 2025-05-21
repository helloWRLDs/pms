-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS "Project" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title" VARCHAR(255),
    "status" VARCHAR(60),
    "description" TEXT,
    "codename" VARCHAR(255),
    "code_prefix" VARCHAR(60),
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "company_id" UUID DEFAULT NULL,
    "progress" INT
);

CREATE TABLE IF NOT EXISTS "Sprint" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title" VARCHAR(255) NOT NULL,
    "description" TEXT,
    "start_date" DATE,
    "end_date" DATE,
    "project_id" UUID DEFAULT NULL,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("project_id") REFERENCES "Project"("id")
);

CREATE TABLE IF NOT EXISTS "Task" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "title" VARCHAR(255),
    "body" TEXT,
    "code" VARCHAR(40),
    "project_id" UUID DEFAULT NULL,
    "sprint_id" UUID DEFAULT NULL,
    "status" TEXT,
    "priority" INTEGER DEFAULT 1,
    "due_date" DATE,
    "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY ("project_id") REFERENCES "Project"("id"),
    FOREIGN KEY ("sprint_id") REFERENCES "Sprint"("id")
);

CREATE TABLE IF NOT EXISTS "SubTask" (
    "parent_id" UUID,
    "child_id" UUID,
    PRIMARY KEY("parent_id", "child_id"),
    FOREIGN KEY ("parent_id") REFERENCES "Task"("id"),
    FOREIGN KEY ("child_id") REFERENCES "Task"("id")
);

CREATE TABLE IF NOT EXISTS "TaskAssignment" (
    "user_id" UUID,
    "task_id" UUID,
    PRIMARY KEY("user_id", "task_id"),
    FOREIGN KEY ("task_id") REFERENCES "Task"("id")
);

CREATE TABLE IF NOT EXISTS "TaskComment" (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "task_id" UUID NOT NULL,
    "user_id" UUID NOT NULL,
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