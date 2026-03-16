// Package repository provides generic repository interfaces for DDD applications.
package repository

import "context"

type Reader[ID any, T any] interface {
	FindByID(ctx context.Context, id ID) (T, error)
}

type Writer[T any] interface {
	Create(ctx context.Context, entity T) error
	Update(ctx context.Context, entity T) error
	Delete(ctx context.Context, entity T) error
}

type Repository[ID any, T any] interface {
	Reader[ID, T]
	Writer[T]
}
