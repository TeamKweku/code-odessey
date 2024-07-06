-- name: CreateFavorite :one
INSERT INTO favorites (
  blog_id,
  user_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetFavorite :one
SELECT * FROM favorites
WHERE id = $1 LIMIT 1;

-- name: ListFavoritesByBlog :many
SELECT * FROM favorites
WHERE blog_id = $1
ORDER BY created_at DESC
LIMIT $2
OFFSET $3;

-- name: DeleteFavorite :exec
DELETE FROM favorites
WHERE id = $1;

-- name: DeleteFavoritesByBlog :execresult
DELETE FROM favorites
WHERE blog_id = $1;