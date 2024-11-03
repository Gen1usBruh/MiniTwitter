// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: retweet.sql

package postgresdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createRetweet = `-- name: CreateRetweet :one
INSERT INTO retweet (
    user_id, parent_tweet_id, parent_retweet_id, is_quote, quote
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, user_id, parent_tweet_id, parent_retweet_id, is_quote, quote, created_at, updated_at
`

type CreateRetweetParams struct {
	UserID          int32       `json:"user_id"`
	ParentTweetID   pgtype.Int4 `json:"parent_tweet_id"`
	ParentRetweetID pgtype.Int4 `json:"parent_retweet_id"`
	IsQuote         bool        `json:"is_quote"`
	Quote           pgtype.Text `json:"quote"`
}

func (q *Queries) CreateRetweet(ctx context.Context, arg CreateRetweetParams) (Retweet, error) {
	row := q.db.QueryRow(ctx, createRetweet,
		arg.UserID,
		arg.ParentTweetID,
		arg.ParentRetweetID,
		arg.IsQuote,
		arg.Quote,
	)
	var i Retweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ParentTweetID,
		&i.ParentRetweetID,
		&i.IsQuote,
		&i.Quote,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteRetweet = `-- name: DeleteRetweet :exec
DELETE FROM retweet 
WHERE id = $1
`

func (q *Queries) DeleteRetweet(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteRetweet, id)
	return err
}

const selectRetweet = `-- name: SelectRetweet :one
SELECT id, user_id, parent_tweet_id, parent_retweet_id, is_quote, quote, created_at, updated_at FROM retweet
WHERE id = $1 LIMIT 1
`

func (q *Queries) SelectRetweet(ctx context.Context, id int32) (Retweet, error) {
	row := q.db.QueryRow(ctx, selectRetweet, id)
	var i Retweet
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ParentTweetID,
		&i.ParentRetweetID,
		&i.IsQuote,
		&i.Quote,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const selectRetweetsOfRetweet = `-- name: SelectRetweetsOfRetweet :many
SELECT r.id AS retweet_id, r.user_id, u.id, r.is_quote, r.quote, r.created_at
FROM retweet r
JOIN users u ON r.user_id = u.id
WHERE r.parent_retweet_id = $1
ORDER BY r.created_at DESC
`

type SelectRetweetsOfRetweetRow struct {
	RetweetID int32              `json:"retweet_id"`
	UserID    int32              `json:"user_id"`
	ID        int32              `json:"id"`
	IsQuote   bool               `json:"is_quote"`
	Quote     pgtype.Text        `json:"quote"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) SelectRetweetsOfRetweet(ctx context.Context, parentRetweetID pgtype.Int4) ([]SelectRetweetsOfRetweetRow, error) {
	rows, err := q.db.Query(ctx, selectRetweetsOfRetweet, parentRetweetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectRetweetsOfRetweetRow{}
	for rows.Next() {
		var i SelectRetweetsOfRetweetRow
		if err := rows.Scan(
			&i.RetweetID,
			&i.UserID,
			&i.ID,
			&i.IsQuote,
			&i.Quote,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const selectRetweetsOfTweet = `-- name: SelectRetweetsOfTweet :many
SELECT r.id AS retweet_id, r.user_id, u.id, r.is_quote, r.quote, r.created_at
FROM retweet r
JOIN users u ON r.user_id = u.id
WHERE r.parent_tweet_id = $1
ORDER BY r.created_at DESC
`

type SelectRetweetsOfTweetRow struct {
	RetweetID int32              `json:"retweet_id"`
	UserID    int32              `json:"user_id"`
	ID        int32              `json:"id"`
	IsQuote   bool               `json:"is_quote"`
	Quote     pgtype.Text        `json:"quote"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

func (q *Queries) SelectRetweetsOfTweet(ctx context.Context, parentTweetID pgtype.Int4) ([]SelectRetweetsOfTweetRow, error) {
	rows, err := q.db.Query(ctx, selectRetweetsOfTweet, parentTweetID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectRetweetsOfTweetRow{}
	for rows.Next() {
		var i SelectRetweetsOfTweetRow
		if err := rows.Scan(
			&i.RetweetID,
			&i.UserID,
			&i.ID,
			&i.IsQuote,
			&i.Quote,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
