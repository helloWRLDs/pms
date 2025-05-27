-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "User" (
	"id" UUID UNIQUE DEFAULT uuid_generate_v4(),
	"email" VARCHAR(255) NOT NULL UNIQUE,
	"password" VARCHAR(60),
	"first_name" VARCHAR(255) NOT NULL,
	"last_name" VARCHAR(255) NOT NULL,
	"phone" VARCHAR(20),
	"bio" TEXT,
	"avatar_img" BYTEA,
	"avatar_url" TEXT,
	"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "Company" (
	"id" UUID UNIQUE DEFAULT uuid_generate_v4(),
	"name" VARCHAR(255) UNIQUE,
	"codename" VARCHAR(255),
	"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("id")
);

CREATE TABLE IF NOT EXISTS "Role" (
	"name" VARCHAR(60) NOT NULL UNIQUE,
	"permissions" TEXT[] NOT NULL,
	"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	"company_id" UUID DEFAULT NULL,
	PRIMARY KEY("name"),
	FOREIGN KEY("company_id") REFERENCES "Company"("id")
	ON UPDATE NO ACTION ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS "Participant" (
	"id" UUID UNIQUE DEFAULT uuid_generate_v4(),
	"user_id" UUID NOT NULL,
	"company_id" UUID NOT NULL,
	"role" VARCHAR NOT NULL,
	"created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	"updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	PRIMARY KEY("id"),
	FOREIGN KEY ("user_id") REFERENCES "User"("id")
	ON UPDATE NO ACTION ON DELETE CASCADE,
	FOREIGN KEY ("company_id") REFERENCES "Company"("id")
	ON UPDATE NO ACTION ON DELETE CASCADE,
	FOREIGN KEY ("role") REFERENCES "Role"("name")
	ON UPDATE NO ACTION ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "Participant";
DROP TABLE "Company";
DROP TABLE "User";
-- +goose StatementEnd
