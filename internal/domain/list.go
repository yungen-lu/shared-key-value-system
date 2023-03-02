package domain

type List struct {
	ID         int32
	PageCount  uint16
	NextPageID int32
}

type ListRepo interface {
	GetHead(id int32) (List, error)
}
