// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: post.sql

package db

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
  title, body
) VALUES (
  $1, $2
)
RETURNING id, title, body, user_id, created_at
`

type CreatePostParams struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.Title, arg.Body)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id, title, body, user_id, created_at FROM posts
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id int32) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id, title, body, user_id, created_at FROM posts
ORDER BY created_at
`

func (q *Queries) ListPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Body,
			&i.UserID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET title = $2,
body = $3
WHERE id = $1
RETURNING id, title, body, user_id, created_at
`

type UpdatePostParams struct {
	ID    int32  `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePost, arg.ID, arg.Title, arg.Body)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.UserID,
		&i.CreatedAt,
	)
	return i, err
}
