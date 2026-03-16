// Package domain contains the core business logic and entities of the application.
package domain

type Entity[ID any] struct {
	ID ID
}
