-- name: CreateTweet :one
INSERT INTO tweet (
    user_id, content, media 
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: SelectTweet :one
SELECT * FROM tweet 
WHERE id = $1 LIMIT 1;

-- name: DeleteTweet :exec
DELETE FROM tweet 
WHERE id = $1;