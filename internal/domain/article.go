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
