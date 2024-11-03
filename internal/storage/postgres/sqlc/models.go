// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package postgresdb

import (
	"database/sql/driver"
	"fmt"

	"github.com/jackc/pgx/v5/pgtype"
)

type TweetType string

const (
	TweetTypeTweet   TweetType = "tweet"
	TweetTypeRetweet TweetType = "retweet"
)

func (e *TweetType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = TweetType(s)
	case string:
		*e = TweetType(s)
	default:
		return fmt.Errorf("unsupported scan type for TweetType: %T", src)
	}
	return nil
}

type NullTweetType struct {
	TweetType TweetType `json:"tweet_type"`
	Valid     bool      `json:"valid"` // Valid is true if TweetType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullTweetType) Scan(value interface{}) error {
	if value == nil {
		ns.TweetType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.TweetType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullTweetType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.TweetType), nil
}

type Comment struct {
	ID              int32              `json:"id"`
	UserID          int32              `json:"user_id"`
	TweetID         pgtype.Int4        `json:"tweet_id"`
	RetweetID       pgtype.Int4        `json:"retweet_id"`
	ParentCommentID pgtype.Int4        `json:"parent_comment_id"`
	PostType        TweetType          `json:"post_type"`
	Content         string             `json:"content"`
	Media           []byte             `json:"media"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
}

type Follow struct {
	FollowerID  int32              `json:"follower_id"`
	FollowingID int32              `json:"following_id"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	UpdatedAt   pgtype.Timestamptz `json:"updated_at"`
}

type LikeComment struct {
	ID        int32              `json:"id"`
	UserID    int32              `json:"user_id"`
	CommentID int32              `json:"comment_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type LikeTweet struct {
	ID        int32              `json:"id"`
	UserID    int32              `json:"user_id"`
	TweetID   pgtype.Int4        `json:"tweet_id"`
	RetweetID pgtype.Int4        `json:"retweet_id"`
	PostType  TweetType          `json:"post_type"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type RefreshToken struct {
	ID        int32              `json:"id"`
	UserID    int32              `json:"user_id"`
	Token     string             `json:"token"`
	ExpiresAt pgtype.Timestamptz `json:"expires_at"`
}

type Retweet struct {
	ID              int32              `json:"id"`
	UserID          int32              `json:"user_id"`
	ParentTweetID   pgtype.Int4        `json:"parent_tweet_id"`
	ParentRetweetID pgtype.Int4        `json:"parent_retweet_id"`
	IsQuote         bool               `json:"is_quote"`
	Quote           pgtype.Text        `json:"quote"`
	CreatedAt       pgtype.Timestamptz `json:"created_at"`
	UpdatedAt       pgtype.Timestamptz `json:"updated_at"`
}

type Tweet struct {
	ID        int32              `json:"id"`
	UserID    int32              `json:"user_id"`
	Content   string             `json:"content"`
	Media     []byte             `json:"media"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}

type User struct {
	ID        int32              `json:"id"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Bio       pgtype.Text        `json:"bio"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	UpdatedAt pgtype.Timestamptz `json:"updated_at"`
}
