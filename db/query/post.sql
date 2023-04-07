-- name: CreatePost :one
INSERT INTO posts (
  title, body, user_id
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1
LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY created_at
LIMIT $1
OFFSET $2;

-- name: UpdatePost :one
UPDATE posts
SET title = $2,
body = $3,
user_id = $4
WHERE id = $1
RETURNING *;


-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1;
