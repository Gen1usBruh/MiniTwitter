-- name: CreateFollower :one
INSERT INTO follows (
    follower_id, following_id
) VALUES (
    $1, $2
) ON CONFLICT DO NOTHING
RETURNING *;

-- name: DeleteFollower :exec
DELETE FROM follows 
WHERE follower_id = $1 AND following_id = $2;