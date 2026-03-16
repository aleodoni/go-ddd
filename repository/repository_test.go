// Package repository provides generic repository interfaces for DDD applications.
package repository_test

import (
	"context"
	"testing"

	"github.com/aleodoni/go-ddd/repository"
)

// --- fakes ---

type fakeEntity struct {
	ID string
}

type fakeReader struct{}

func (f *fakeReader) FindByID(ctx context.Context, id string) (*fakeEntity, error) {
	return nil, nil
}

type fakeWriter struct{}

func (f *fakeWriter) Create(ctx context.Context, entity *fakeEntity) error { return nil }
func (f *fakeWriter) Update(ctx context.Context, entity *fakeEntity) error { return nil }
func (f *fakeWriter) Delete(ctx context.Context, entity *fakeEntity) error { return nil }

type fakeRepository struct {
	fakeReader
	fakeWriter
}

// --- verificações de interface em tempo de compilação ---

func TestRepository_ImplementaReader(t *testing.T) {
	var _ repository.Reader[string, *fakeEntity] = &fakeReader{}
}

func TestRepository_ImplementaWriter(t *testing.T) {
	var _ repository.Writer[*fakeEntity] = &fakeWriter{}
}

func TestRepository_ImplementaRepository(t *testing.T) {
	var _ repository.Repository[string, *fakeEntity] = &fakeRepository{}
}
