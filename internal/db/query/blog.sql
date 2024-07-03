-- name: CreateBlog :one
INSERT INTO blogs (
  title,
  slug,
  description,
  body,
  banner_image
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetBlog :one
SELECT * FROM blogs
WHERE id = $1 LIMIT 1;


-- name: GetBlogBySlug :one
SELECT * FROM blogs
WHERE slug = $1 LIMIT 1;

-- name: ListBlogs :many
SELECT * FROM blogs
ORDER BY created_at DESC
LIMIT $1
OFFSET $2;


-- name: UpdateBlog :one
UPDATE blogs
SET
  title = $2,
  slug = $3,
  description = $4,
  body = $5,
  banner_image = $6,
  updated_at = $7
WHERE id = $1
RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1;