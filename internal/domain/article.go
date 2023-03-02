package domain

import "time"

type Article struct {
	ID        int32
	Title     string
	Content   string
	Author    User
	UpdatedAt time.Time
	CreatedAt time.Time
	Topic     Topic
	Tags      string // TODO
}

type ArticleRepo interface {
	GetByID(id int32) (Article, error)
	GetByTitle(title string) (Article, error)
	Update(article *Article) error
	Store(article *Article) error
	DeleteByID(id int32) error
}
