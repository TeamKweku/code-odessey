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
  title = COALESCE(sqlc.narg(title), title),
  slug = COALESCE(sqlc.narg(slug), slug),
  description = COALESCE(sqlc.narg(description), description),
  body = COALESCE(sqlc.narg(body), body),
  banner_image = COALESCE(sqlc.narg(banner_image), banner_image)
WHERE 
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteBlog :exec
DELETE FROM blogs
WHERE id = $1;