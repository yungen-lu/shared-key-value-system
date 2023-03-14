-- name: GetUserByID :one 
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at;

-- name: GetTopicByID :one
SELECT * FROM topics
WHERE id = $1 LIMIT 1;

-- name: ListTopics :many
SELECT * FROM topics
ORDER BY created_at;

-- name: GetPageByKey :one
SELECT * FROM pages
WHERE key = $1 LIMIT 1;

-- name: GetPageByID :one
SELECT * FROM pages
WHERE id = $1 LIMIT 1;

-- name: UpdatePageByKey :one
UPDATE pages
SET
  key = sqlc.arg('key'),
  next_id = sqlc.arg('next_id')
WHERE key = sqlc.arg('oldkey')
RETURNING *;


-- name: ListPages :many
SELECT * FROM pages
ORDER BY created_at;

-- name: CreatePage :one
INSERT INTO pages (key, next_id) VALUES ($1, $2) RETURNING *;

-- name: GetListByKey :one
SELECT * FROM lists
WHERE key = $1 LIMIT 1;

-- name: GetListByID :one
SELECT * FROM lists
WHERE id = $1 LIMIT 1;

-- name: UpdateListByKey :one
UPDATE lists
SET
  key = sqlc.arg('key'),
  next_page_id = sqlc.arg('next_page_id')
WHERE key = sqlc.arg('oldkey')
RETURNING *;

-- name: ListLists :many
SELECT * FROM lists
ORDER BY created_at;

-- name: CreateList :one
INSERT INTO lists (key, page_count, next_page_id) VALUES ($1, $2, $3) RETURNING *;

-- name: GetArticleByID :one
SELECT * FROM articles
WHERE id = $1 LIMIT 1;

-- name: ListArticles :many
SELECT * FROM articles
ORDER BY created_at;
