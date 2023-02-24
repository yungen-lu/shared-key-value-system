package domain

type Page struct {
	ID         uint64 // key
	Articles   []Article
	NextPageID uint64
}
