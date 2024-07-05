-- name: CreateComment :one
INSERT INTO comments (
  blog_id,
  body
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetComment :one
SELECT * FROM comments
WHERE id = $1 LIMIT 1;

-- name: ListCommentsByBlog :many
SELECT * FROM comments
WHERE blog_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: UpdateComment :one
UPDATE comments
SET
  body = $2
WHERE id = $1
RETURNING *;

-- name: UpdateCommentByBlogID :one
UPDATE comments
SET body = $3
WHERE id = $1 AND blog_id = $2
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1;

-- name: DeleteCommentByBlogID :exec
DELETE FROM comments
WHERE id = $1 AND blog_id = $2;

-- name: DeleteCommentsByBlog :execresult
DELETE FROM comments
WHERE blog_id = $1;