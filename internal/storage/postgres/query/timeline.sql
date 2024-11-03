-- name: SelectTimeline :many
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
LIMIT $2 OFFSET $3;