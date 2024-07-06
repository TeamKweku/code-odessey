-- Remove foreign key constraints added in the latest migration
ALTER TABLE IF EXISTS blogs DROP CONSTRAINT IF EXISTS blogs_author_id_fkey;
ALTER TABLE IF EXISTS comments DROP CONSTRAINT IF EXISTS comments_user_id_fkey;
ALTER TABLE IF EXISTS favorites DROP CONSTRAINT IF EXISTS favorites_user_id_fkey;

-- Remove columns added to existing tables
ALTER TABLE IF EXISTS blogs DROP COLUMN IF EXISTS author_id;
ALTER TABLE IF EXISTS comments DROP COLUMN IF EXISTS user_id;
ALTER TABLE IF EXISTS favorites DROP COLUMN IF EXISTS user_id;

-- Remove indexes created in the latest migration
DROP INDEX IF EXISTS users_email_idx;
DROP INDEX IF EXISTS users_username_idx;
DROP INDEX IF EXISTS blogs_author_id_created_at_idx;
DROP INDEX IF EXISTS blogs_slug_idx;
DROP INDEX IF EXISTS blogs_title_idx;
DROP INDEX IF EXISTS blogs_created_at_idx;
DROP INDEX IF EXISTS comments_blog_id_created_at_idx;
DROP INDEX IF EXISTS favorites_blog_id_user_id_idx;

-- Remove trigger for users table
DROP TRIGGER IF EXISTS update_users_updated_at ON users;

-- Drop the users table
DROP TABLE IF EXISTS users;