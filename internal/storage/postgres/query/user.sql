-- name: CreateUser :one
INSERT INTO users (
    username, email, password, bio 
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateUserBio :exec
UPDATE users 
    set bio = $2 
WHERE id = $1;

-- name: UpdateUserName :exec
UPDATE users 
    set username = $2 
WHERE id = $1;

-- name: SelectUserCred :one
SELECT username, email, password FROM users 
WHERE id = $1 LIMIT 1;

-- name: SelectUserData :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;

-- name: SelectUserFollowers :many
SELECT u.id, u.username, u.bio, f.created_at AS follow_date
FROM follows f
JOIN users u ON f.follower_id = u.id
WHERE f.following_id = $1
ORDER BY f.created_at DESC;

-- name: SelectUserFollowing :many
SELECT u.id, u.username, u.bio, f.created_at AS follow_date
FROM follows f
JOIN users u ON f.following_id = u.id
WHERE f.follower_id = $1
ORDER BY f.created_at DESC;


