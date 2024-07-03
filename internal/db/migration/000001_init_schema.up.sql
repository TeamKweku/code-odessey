CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "blogs" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "title" varchar(250) NOT NULL ,
  "slug" varchar(255) UNIQUE NOT NULL,
  "description" varchar(255) NOT NULL,
  "body" text NOT NULL,
  "banner_image" varchar(255) NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "comments" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "blog_id" UUID NOT NULL,
  "body" text NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE TABLE "favorites" (
  "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  "blog_id" UUID NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

CREATE INDEX ON "comments" ("blog_id");

CREATE INDEX ON "favorites" ("blog_id");

ALTER TABLE "comments" ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id");

ALTER TABLE "favorites" ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id");