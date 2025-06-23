-- name: Createfeed :one
INSERT INTO feeds (id, created_at, updated_at, name, user_id, url)
  VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
  *;

-- name: GetFeeds :many
SELECT
  feeds.id,
  feeds.created_at,
  feeds.updated_at,
  feeds.name,
  feeds.user_id,
  feeds.url,
  "users".name AS user_name
FROM
  feeds
  JOIN "users" ON feeds.user_id = "users".id
ORDER BY
  feeds.created_at DESC;
