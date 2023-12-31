// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.0
// source: posts.sql

package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6,
  $7,
  $8
  )
RETURNING id, created_at, updated_at, title, url, description, published_at, feed_id
`

type CreatePostParams struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt sql.NullTime
	FeedID      uuid.UUID
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.Url,
		arg.Description,
		arg.PublishedAt,
		arg.FeedID,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
	)
	return i, err
}

const getPostByUser = `-- name: GetPostByUser :one
SELECT posts.id, posts.created_at, posts.updated_at, title, url, description, published_at, posts.feed_id, feed_follows.id, feed_follows.feed_id, user_id, feed_follows.created_at, feed_follows.updated_at FROM posts
JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_at DESC
LIMIT $2
`

type GetPostByUserParams struct {
	UserID uuid.UUID
	Limit  int32
}

type GetPostByUserRow struct {
	ID          uuid.UUID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description string
	PublishedAt sql.NullTime
	FeedID      uuid.UUID
	ID_2        uuid.UUID
	FeedID_2    uuid.UUID
	UserID      uuid.UUID
	CreatedAt_2 time.Time
	UpdatedAt_2 time.Time
}

func (q *Queries) GetPostByUser(ctx context.Context, arg GetPostByUserParams) (GetPostByUserRow, error) {
	row := q.db.QueryRowContext(ctx, getPostByUser, arg.UserID, arg.Limit)
	var i GetPostByUserRow
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.Url,
		&i.Description,
		&i.PublishedAt,
		&i.FeedID,
		&i.ID_2,
		&i.FeedID_2,
		&i.UserID,
		&i.CreatedAt_2,
		&i.UpdatedAt_2,
	)
	return i, err
}
