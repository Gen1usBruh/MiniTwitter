-- name: CreateComment :one
INSERT INTO comment (
    user_id, tweet_id, retweet_id, parent_comment_id, post_type, content, media
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: SelectComment :many
SELECT c.id, c.user_id, u.username, c.content, c.media, c.created_at
FROM comment c
JOIN users u ON c.user_id = u.id
WHERE (
    (c.tweet_id = $1 AND $2 = 'tweet') 
OR
    (c.retweet_id = $1 AND $2 = 'retweet')
) AND 
    ($1 IS NOT NULL AND $2 <> '')
ORDER BY c.created_at;

-- name: DeleteComment :exec
DELETE FROM comment
WHERE id = $1 AND user_id = $2;