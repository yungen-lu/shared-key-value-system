package domain

type Page struct {
	ID         uint64 // key
	Articles   []Article
	NextPageID uint64
}

type PageRepo interface {
	GetPage(id uint64) (Page, error)
}
