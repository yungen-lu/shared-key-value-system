package domain

type List struct {
	ID         uint64
	PageCount  uint16
	NextPageID uint64
}

type ListRepo interface {
	GetHead(id uint64) (List, error)
}
