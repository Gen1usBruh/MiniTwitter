// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: timeline.sql

package postgresdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const selectTimeline = `-- name: SelectTimeline :many
(
SELECT 
        t.id AS post_id,
        'tweet' AS post_type,
        t.user_id,
        u.username,
        COALESCE(t.content, '') AS content,
        t.media,
        t.created_at,
        NULL AS quote,
        NULL AS parent_tweet_id,
        NULL AS parent_retweet_id
    FROM tweet t
    JOIN follows f ON f.following_id = t.user_id
    JOIN users u ON t.user_id = u.id
    WHERE f.follower_id = $1
)
UNION ALL
(
    SELECT 
        rt.id AS post_id,
        'retweet' AS post_type,
        rt.user_id,
        u.username,
        COALESCE(NULL, '') AS content,
        NULL AS media,
        rt.created_at,
        rt.quote,
        rt.parent_tweet_id,
        rt.parent_retweet_id
    FROM retweet rt
    JOIN follows f ON f.following_id = rt.user_id
    JOIN users u ON rt.user_id = u.id
    WHERE f.follower_id = $1
)
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type SelectTimelineParams struct {
	FollowerID int32 `json:"follower_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

type SelectTimelineRow struct {
	PostID          int32              `json:"post_id"`
	PostType        string             `json:"post_type"`
	UserID          int32              `json:"user_id"`
	Username        string             `json:"username"`
	Content         string             `json:"content"`
	Media           []byte             `json:"media"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	Quote           interface{}        `json:"quote"`
	ParentTweetID   interface{}        `json:"parent_tweet_id"`
	ParentRetweetID interface{}        `json:"parent_retweet_id"`
}

func (q *Queries) SelectTimeline(ctx context.Context, arg SelectTimelineParams) ([]SelectTimelineRow, error) {
	rows, err := q.db.Query(ctx, selectTimeline, arg.FollowerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []SelectTimelineRow{}
	for rows.Next() {
		var i SelectTimelineRow
		if err := rows.Scan(
			&i.PostID,
			&i.PostType,
			&i.UserID,
			&i.Username,
			&i.Content,
			&i.Media,
			&i.CreatedAt,
			&i.Quote,
			&i.ParentTweetID,
			&i.ParentRetweetID,
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
