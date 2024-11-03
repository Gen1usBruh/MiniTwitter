-- name: CreateRetweet :one
INSERT INTO retweet (
    user_id, parent_tweet_id, parent_retweet_id, is_quote, quote
) VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: SelectRetweet :one
SELECT * FROM retweet
WHERE id = $1 LIMIT 1;

-- name: DeleteRetweet :exec
DELETE FROM retweet 
WHERE id = $1;

-- name: SelectRetweetsOfTweet :many
SELECT r.id AS retweet_id, r.user_id, u.id, r.is_quote, r.quote, r.created_at
FROM retweet r
JOIN users u ON r.user_id = u.id
WHERE r.parent_tweet_id = $1
ORDER BY r.created_at DESC;

-- name: SelectRetweetsOfRetweet :many
SELECT r.id AS retweet_id, r.user_id, u.id, r.is_quote, r.quote, r.created_at
FROM retweet r
JOIN users u ON r.user_id = u.id
WHERE r.parent_retweet_id = $1
ORDER BY r.created_at DESC;