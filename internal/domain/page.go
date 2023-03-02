package domain

type Page struct {
	ID         int32 // key
	Articles   []Article
	NextPageID int32
}

type PageRepo interface {
	GetPage(id int32) (Page, error)
}
