-- name: Createfeed :one
INSERT INTO feed (id, created_at, updated_at, name, user_id, url)
  VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
  *;
