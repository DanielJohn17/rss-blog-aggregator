-- name: CreateFeedFollow :one
WITH inserted AS (
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
  RETURNING
    *)
  SELECT
    inserted.*,
    users.name AS user_name,
    feeds.name AS feed_name
  FROM
    inserted
    JOIN users ON inserted.user_id = users.id
    JOIN feeds ON inserted.feed_id = feeds.id;

-- name: GetFeedFollowsForUser :many
SELECT
  ff.*,
  u.name AS user_name,
  f.name AS feed_name
FROM
  feed_follows ff
  JOIN users u ON ff.user_is = u.id
  JOIN feeds f ON ff.feed_id = f.id
WHERE
  ff.user_id = $1;
