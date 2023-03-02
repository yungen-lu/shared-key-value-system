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

-- name: GetPageByID :one
SELECT * FROM pages
WHERE id = $1 LIMIT 1;

-- name: ListPages :many
SELECT * FROM pages
ORDER BY created_at;


-- name: GetListByID :one
SELECT * FROM lists
WHERE id = $1 LIMIT 1;

-- name: ListLists :many
SELECT * FROM lists
ORDER BY created_at;

-- name: GetArticleByID :one
SELECT * FROM articles
WHERE id = $1 LIMIT 1;

-- name: ListArticles :many
SELECT * FROM articles
ORDER BY created_at;
