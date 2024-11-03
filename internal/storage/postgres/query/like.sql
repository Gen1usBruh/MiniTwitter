-- name: CreateLikeTweet :one
INSERT INTO like_tweet (
    user_id, tweet_id, retweet_id, post_type
)
VALUES (
    $1, $2, $3, $4
)
ON CONFLICT DO NOTHING
RETURNING *;

-- name: SelectLikeTweet :many
SELECT lt.id, lt.user_id, u.username, lt.created_at
FROM like_tweet lt
JOIN users u ON lt.user_id = u.id
WHERE ( 
    ($1 = 'tweet' AND lt.tweet_id = $2) 
OR
    ($1 = 'retweet' AND lt.retweet_id = $2)
)
ORDER BY lt.created_at DESC;

-- name: DeleteLikeTweet :exec 
DELETE FROM like_tweet
WHERE user_id = $1 AND 
(   ($2 = 'tweet' AND tweet_id = $3) 
OR
    ($2 = 'retweet' AND retweet_id = $3)
);

-- name: CreateLikeComment :one
INSERT INTO like_comment (
    user_id, comment_id
)
VALUES (
    $1, $2
)
ON CONFLICT DO NOTHING
RETURNING *;

-- name: SelectLikeComment :many
SELECT lc.id, lc.user_id, u.username, lc.created_at
FROM like_comment lc
JOIN users u ON lc.user_id = u.id
WHERE lc.comment_id = $1
ORDER BY lc.created_at DESC;

-- name: DeleteLikeComment :exec 
DELETE FROM like_comment
WHERE user_id = $1 AND comment_id = $2;