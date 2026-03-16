# go-ddd

A lightweight Go library with building blocks for applications using **Domain-Driven Design (DDD)**.

## Requirements

- Go 1.21+

## Installation

```bash
go get github.com/aleodoni/go-ddd
```

## Packages

### `domain`

Core building blocks for your domain model.

#### `Entity[ID]`

Base struct for domain entities. Uses generics so the ID can be of any type.

```go
type Usuario struct {
    domain.Entity[string]
    Nome  string
    Email string
}

func NewUsuario(nome, email string) *Usuario {
    return &Usuario{
        Entity: domain.Entity[string]{ID: cuid2.Generate()},
        Nome:   nome,
        Email:  email,
    }
}
```

#### `AggregateRoot[ID]`

Extends `Entity` with domain event support. Use it as the root of your aggregates.

```go
type Pedido struct {
    domain.AggregateRoot[string]
    Items []ItemPedido
}

func NewPedido() *Pedido {
    p := &Pedido{
        AggregateRoot: domain.NewAggregateRoot[string](cuid2.Generate()),
    }
    p.RaiseEvent(PedidoCriadoEvent{PedidoID: p.ID})
    return p
}

// later, dispatch the events
events := pedido.PullEvents()
```

| Method | Description |
|--------|-------------|
| `NewAggregateRoot[ID](id)` | Creates a new AggregateRoot with the given ID |
| `RaiseEvent(event)` | Registers a domain event |
| `PullEvents()` | Returns all registered events and clears the list |

#### `DomainEvent`

Interface that all domain events must implement.

```go
type PedidoCriadoEvent struct {
    PedidoID   string
    occurredAt time.Time
}

func (e PedidoCriadoEvent) EventName() string     { return "pedido.criado" }
func (e PedidoCriadoEvent) OccurredAt() time.Time { return e.occurredAt }
```

#### `ValueObject`

Interface for value objects. Implementors must define equality by value, not identity.

```go
type Email struct {
    value string
}

func NewEmail(value string) (Email, error) {
    if !strings.Contains(value, "@") {
        return Email{}, errors.New("email inválido")
    }
    return Email{value: value}, nil
}

func (e Email) Equals(other domain.ValueObject) bool {
    o, ok := other.(Email)
    return ok && e.value == o.value
}

func (e Email) String() string { return e.value }
```

---

### `errors`

Standardized domain errors.

#### Sentinel errors

```go
import dderr "github.com/aleodoni/go-ddd/errors"

return dderr.ErrNotFound
return dderr.ErrAlreadyExists
return dderr.ErrForbidden
return dderr.ErrUnauthorized
return dderr.ErrInvalidInput
return dderr.ErrConflict
return dderr.ErrInternal
```

#### `DomainError`

For errors that require additional context.

```go
// with code and message
return dderr.New("USER_INACTIVE", "usuário inativo", nil)

// wrapping a cause
return dderr.New("DB_ERROR", "erro ao buscar usuário", err)
```

```go
type DomainError struct {
    Code    string
    Message string
    Cause   error
}
```

---

### `repository`

Generic repository interfaces following the **Interface Segregation Principle**.

```go
import "github.com/aleodoni/go-ddd/repository"

// read-only
type Reader[ID any, T any] interface {
    FindByID(ctx context.Context, id ID) (T, error)
}

// write-only
type Writer[T any] interface {
    Create(ctx context.Context, entity T) error
    Update(ctx context.Context, entity T) error
    Delete(ctx context.Context, entity T) error
}

// full CRUD
type Repository[ID any, T any] interface {
    Reader[ID, T]
    Writer[T]
}
```

Use `Reader` or `Writer` when a use case only needs part of the contract:

```go
// a report service only needs to read
type RelatorioService struct {
    repo repository.Reader[string, *Usuario]
}

// an import service only needs to write
type ImportacaoService struct {
    repo repository.Writer[*Usuario]
}
```

Extend with domain-specific methods:

```go
type UsuarioRepository interface {
    repository.Repository[string, *Usuario]
    FindByEmail(ctx context.Context, email string) (*Usuario, error)
    FindByUsername(ctx context.Context, username string) (*Usuario, error)
}
```

---

## Running tests

```bash
go test ./...
```

## License

MIT
