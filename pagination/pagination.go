// Package pagination provides pagination utilities for DDD applications.
package pagination

type Params struct {
	Page  int
	Limit int
}

func NewParams(page, limit int) Params {
	p := Params{Page: page, Limit: limit}
	p.Normalize()
	return p
}

func (p *Params) Normalize() {
	if p.Page < 1 {
		p.Page = 1
	}
	if p.Limit < 1 || p.Limit > 100 {
		p.Limit = 20
	}
}

func (p Params) Offset() int {
	return (p.Page - 1) * p.Limit
}

type PagedResult[T any] struct {
	Items []T
	Total int64
	Page  int
	Limit int
}
