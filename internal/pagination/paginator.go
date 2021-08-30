package pagination

import "math"

type Paginator interface {
	GetTotal() uint64
	GetPerPage() uint64
	GetCurrentPage() uint64
	GetLastPage() uint64
}

type paginator struct {
	total       uint64
	perPage     uint64
	currentPage uint64
}

func New(page, perPage, total uint64) Paginator {
	return &paginator{
		total:       total,
		perPage:     perPage,
		currentPage: page,
	}
}

func (p *paginator) GetTotal() uint64 {
	return p.total
}

func (p *paginator) GetPerPage() uint64 {
	return p.perPage
}

func (p *paginator) GetCurrentPage() uint64 {
	return p.currentPage
}

func (p *paginator) GetLastPage() uint64 {
	return uint64(math.Ceil(float64(p.total) / float64(p.perPage)))
}
