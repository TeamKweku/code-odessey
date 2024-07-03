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

-- Function to update the updated_at column
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
END;
$$ language 'plpgsql';

-- Trigger for blogs table
CREATE TRIGGER update_blogs_updated_at
BEFORE UPDATE ON blogs
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Trigger for comments table
CREATE TRIGGER update_comments_updated_at
BEFORE UPDATE ON comments
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

-- Trigger for favorites table
CREATE TRIGGER update_favorites_updated_at
BEFORE UPDATE ON favorites
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();