-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2024-07-17T22:11:51.387Z

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar(20) NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar(250) NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_eamil_verified" bool NOT NULL DEFAULT false,
  "password_changed_at" timestampz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "verify_emails" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "user_id" uuid NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestampz NOT NULL DEFAULT (now()),
  "expires_at" timestamptz NOT NULL DEFAULT (now() + interval '15 minutes')
);

CREATE TABLE "blog" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "author_id" uuid NOT NULL,
  "title" varchar(250) NOT NULL,
  "slug" varchar(255) UNIQUE NOT NULL,
  "description" varchar(255) NOT NULL,
  "body" text NOT NULL,
  "banner_image" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "comment" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "blog_id" uuid NOT NULL,
  "user_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "favorite" (
  "id" uuid PRIMARY KEY,
  "blog_id" uuid,
  "user_id" uuid NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "sessions" (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("email");

CREATE INDEX ON "users" ("username");

CREATE INDEX ON "blog" ("author_id", "created_at");

CREATE INDEX ON "blog" ("slug");

CREATE INDEX ON "blog" ("title");

CREATE INDEX ON "blog" ("created_at");

CREATE INDEX ON "comment" ("blog_id", "created_at");

CREATE INDEX ON "favorite" ("blog_id", "user_id");

CREATE INDEX ON "sessions" ("user_id");

ALTER TABLE "verify_emails" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "blog" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

ALTER TABLE "comment" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "favorite" ADD FOREIGN KEY ("blog_id") REFERENCES "blog" ("id");

ALTER TABLE "favorite" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
