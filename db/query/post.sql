-- name: CreatePost :one
INSERT INTO posts (
  title, body
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1
LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY created_at;

-- name: UpdatePost :one
UPDATE posts
SET title = $2,
body = $3
WHERE id = $1
RETURNING *;


-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;
