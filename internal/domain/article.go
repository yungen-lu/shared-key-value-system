package domain

import "time"

type Article struct {
	ID        uint64
	Title     string
	Content   string
	Author    User
	UpdatedAt time.Time
	CreatedAt time.Time
	Topic     Topic
	Tags      string // TODO
}

type ArticleRepo interface {
	GetByID(id uint64) (Article, error)
	GetByTitle(title string) (Article, error)
	Update(article *Article) error
	Store(article *Article) error
	DeleteByID(id uint64) error
}
