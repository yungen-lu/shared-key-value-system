// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2
// source: query.sql

package pgcodegen

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createList = `-- name: CreateList :one
INSERT INTO lists (id, page_count, next_page_id) VALUES ($1, $2, $3) RETURNING id, page_count, next_page_id, created_at, updated_at
`

type CreateListParams struct {
	ID         int32
	PageCount  int32
	NextPageID pgtype.Int4
}

func (q *Queries) CreateList(ctx context.Context, arg CreateListParams) (List, error) {
	row := q.db.QueryRow(ctx, createList, arg.ID, arg.PageCount, arg.NextPageID)
	var i List
	err := row.Scan(
		&i.ID,
		&i.PageCount,
		&i.NextPageID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createPage = `-- name: CreatePage :one
INSERT INTO pages (next_id) VALUES ($1) RETURNING id, next_id, created_at, updated_at
`

func (q *Queries) CreatePage(ctx context.Context, nextID pgtype.Int4) (Page, error) {
	row := q.db.QueryRow(ctx, createPage, nextID)
	var i Page
	err := row.Scan(
		&i.ID,
		&i.NextID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getArticleByID = `-- name: GetArticleByID :one
SELECT id, title, content, author_id, topic_id, created_at, updated_at FROM articles
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetArticleByID(ctx context.Context, id int32) (Article, error) {
	row := q.db.QueryRow(ctx, getArticleByID, id)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.AuthorID,
		&i.TopicID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getListByID = `-- name: GetListByID :one
SELECT id, page_count, next_page_id, created_at, updated_at FROM lists
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetListByID(ctx context.Context, id int32) (List, error) {
	row := q.db.QueryRow(ctx, getListByID, id)
	var i List
	err := row.Scan(
		&i.ID,
		&i.PageCount,
		&i.NextPageID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPageByID = `-- name: GetPageByID :one
SELECT id, next_id, created_at, updated_at FROM pages
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPageByID(ctx context.Context, id int32) (Page, error) {
	row := q.db.QueryRow(ctx, getPageByID, id)
	var i Page
	err := row.Scan(
		&i.ID,
		&i.NextID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTopicByID = `-- name: GetTopicByID :one
SELECT id, name, description, created_at, updated_at FROM topics
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTopicByID(ctx context.Context, id int32) (Topic, error) {
	row := q.db.QueryRow(ctx, getTopicByID, id)
	var i Topic
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, created_at, updated_at FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listArticles = `-- name: ListArticles :many
SELECT id, title, content, author_id, topic_id, created_at, updated_at FROM articles
ORDER BY created_at
`

func (q *Queries) ListArticles(ctx context.Context) ([]Article, error) {
	rows, err := q.db.Query(ctx, listArticles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Article
	for rows.Next() {
		var i Article
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.AuthorID,
			&i.TopicID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listLists = `-- name: ListLists :many
SELECT id, page_count, next_page_id, created_at, updated_at FROM lists
ORDER BY created_at
`

func (q *Queries) ListLists(ctx context.Context) ([]List, error) {
	rows, err := q.db.Query(ctx, listLists)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []List
	for rows.Next() {
		var i List
		if err := rows.Scan(
			&i.ID,
			&i.PageCount,
			&i.NextPageID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPages = `-- name: ListPages :many
SELECT id, next_id, created_at, updated_at FROM pages
ORDER BY created_at
`

func (q *Queries) ListPages(ctx context.Context) ([]Page, error) {
	rows, err := q.db.Query(ctx, listPages)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Page
	for rows.Next() {
		var i Page
		if err := rows.Scan(
			&i.ID,
			&i.NextID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTopics = `-- name: ListTopics :many
SELECT id, name, description, created_at, updated_at FROM topics
ORDER BY created_at
`

func (q *Queries) ListTopics(ctx context.Context) ([]Topic, error) {
	rows, err := q.db.Query(ctx, listTopics)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Topic
	for rows.Next() {
		var i Topic
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, created_at, updated_at FROM users
ORDER BY created_at
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
