CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "users" (
  "id" uuid PRIMARY KEY DEFAULT (uuid_generate_v4()),
  "username" varchar(20) NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar(250) NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "is_email_verified" bool NOT NULL DEFAULT false,
  "password_changed_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamp NOT NULL DEFAULT (now()),
  "updated_at" timestamp NOT NULL DEFAULT '0001-01-01 00:00:00Z'
);

ALTER TABLE "blogs" ADD COLUMN "author_id" uuid NOT NULL;

ALTER TABLE "comments" ADD COLUMN "user_id" uuid NOT NULL;

ALTER TABLE "favorites" ADD COLUMN "user_id" uuid NOT NULL;

CREATE INDEX ON "users" ("email");
CREATE INDEX ON "users" ("username");
CREATE INDEX ON "blogs" ("author_id", "created_at");
CREATE INDEX ON "blogs" ("slug");
CREATE INDEX ON "blogs" ("title");
CREATE INDEX ON "blogs" ("created_at");
CREATE INDEX ON "comments" ("blog_id", "created_at");
CREATE INDEX ON "favorites" ("blog_id", "user_id");


ALTER TABLE "blogs" ADD FOREIGN KEY ("author_id") REFERENCES "users" ("id");
ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
ALTER TABLE "favorites" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


CREATE TRIGGER update_users_updated_at
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
