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
  next_page_key = $1
WHERE key = $2
RETURNING *;

-- name: DeletePageByKey :execrows
DELETE FROM pages
WHERE key = $1;


-- name: ListPages :many
SELECT * FROM pages
ORDER BY created_at;

-- name: CreatePage :one
INSERT INTO pages (key, next_page_key, list_key) VALUES ($1, $2, $3)
RETURNING *;

-- name: GetListByKey :one
SELECT * FROM lists
WHERE key = $1 LIMIT 1;

-- name: GetListByID :one
SELECT * FROM lists
WHERE id = $1 LIMIT 1;

-- name: UpdateListByKey :one
UPDATE lists
SET
  next_page_key = $1,
  latest_page_key = $2
WHERE key = $3
RETURNING *;

-- name: ListLists :many
SELECT * FROM lists
ORDER BY created_at;

-- name: CreateList :one
INSERT INTO lists (key, next_page_key, latest_page_key) VALUES ($1, $2, $3) RETURNING *;

-- name: DeleteListByKey :execrows
DELETE FROM lists
WHERE key = $1;

-- name: GetArticleByID :one
SELECT * FROM articles
WHERE id = $1 LIMIT 1;

-- name: ListArticles :many
SELECT * FROM articles
ORDER BY created_at;
