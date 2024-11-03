-- name: CreateToken :one
INSERT INTO refresh_tokens (
    user_id, token, expires_at
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: SelectToken :one
SELECT * 
FROM refresh_tokens
WHERE user_id = $1 AND 
    token = $2 AND expires_at > NOW()
LIMIT 1;

-- name: DeleteToken :exec
DELETE FROM refresh_tokens
WHERE user_id = $1 AND token = $2;

