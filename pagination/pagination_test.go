// Package pagination provides pagination utilities for DDD applications.
package pagination_test

import (
	"testing"

	"github.com/aleodoni/go-ddd/pagination"
)

func TestNewParams_ValoresValidos(t *testing.T) {
	p := pagination.NewParams(2, 10)

	if p.Page != 2 {
		t.Errorf("esperava Page = 2, got %d", p.Page)
	}
	if p.Limit != 10 {
		t.Errorf("esperava Limit = 10, got %d", p.Limit)
	}
}

func TestNewParams_PaginaMenorQueUm(t *testing.T) {
	p := pagination.NewParams(0, 10)

	if p.Page != 1 {
		t.Errorf("esperava Page = 1, got %d", p.Page)
	}
}

func TestNewParams_LimitMenorQueUm(t *testing.T) {
	p := pagination.NewParams(1, 0)

	if p.Limit != 20 {
		t.Errorf("esperava Limit = 20, got %d", p.Limit)
	}
}

func TestNewParams_LimitMaiorQueCem(t *testing.T) {
	p := pagination.NewParams(1, 101)

	if p.Limit != 20 {
		t.Errorf("esperava Limit = 20, got %d", p.Limit)
	}
}

func TestParams_Offset_PaginaUm(t *testing.T) {
	p := pagination.NewParams(1, 10)

	if p.Offset() != 0 {
		t.Errorf("esperava Offset = 0, got %d", p.Offset())
	}
}

func TestParams_Offset_PaginaDois(t *testing.T) {
	p := pagination.NewParams(2, 10)

	if p.Offset() != 10 {
		t.Errorf("esperava Offset = 10, got %d", p.Offset())
	}
}

func TestParams_Offset_PaginaTres(t *testing.T) {
	p := pagination.NewParams(3, 10)

	if p.Offset() != 20 {
		t.Errorf("esperava Offset = 20, got %d", p.Offset())
	}
}

func TestPagedResult_Generics(t *testing.T) {
	items := []string{"a", "b", "c"}

	result := pagination.PagedResult[string]{
		Items: items,
		Total: 3,
		Page:  1,
		Limit: 10,
	}

	if len(result.Items) != 3 {
		t.Errorf("esperava 3 items, got %d", len(result.Items))
	}
	if result.Total != 3 {
		t.Errorf("esperava Total = 3, got %d", result.Total)
	}
}
