-- name: CreateFeed :one
INSERT INTO feeds (id, created_at, updated_at, name, url, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: FetchFeeds :many
SELECT 
    feeds.name AS feed_name,
    feeds.url AS feed_url,
    users.name AS user_name
FROM feeds
INNER JOIN users ON feeds.user_id = users.id
ORDER BY feeds.name ASC;

-- name: FetchFeed :one
SELECT
    feeds.name AS feed_name,
    feeds.id AS feed_id
FROM feeds
WHERE feeds.url = $1;

-- name: GetNextFeedToFetch :one
SELECT
    feeds.url AS feeds_url,
    feeds.id AS feeds_id
FROM feeds
ORDER BY feeds.last_fetched_at NULLS FIRST
LIMIT 1;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET 
    updated_at = $1,
    last_fetched_at = $1
WHERE feeds.id = $2;