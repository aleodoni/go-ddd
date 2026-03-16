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
type User struct {
    domain.Entity[string]
    Name  string
    Email string
}

func NewUser(name, email string) *User {
    return &User{
        Entity: domain.Entity[string]{ID: cuid2.Generate()},
        Name:   name,
        Email:  email,
    }
}
```

#### `AggregateRoot[ID]`

Extends `Entity` with domain event support. Use it as the root of your aggregates.

```go
type Order struct {
    domain.AggregateRoot[string]
    Items []OrderItem
}

func NewOrder() *Order {
    o := &Order{
        AggregateRoot: domain.NewAggregateRoot[string](cuid2.Generate()),
    }
    o.RaiseEvent(OrderCreatedEvent{OrderID: o.ID})
    return o
}

// later, dispatch the events
events := order.PullEvents()
```

| Method | Description |
|--------|-------------|
| `NewAggregateRoot[ID](id)` | Creates a new AggregateRoot with the given ID |
| `RaiseEvent(event)` | Registers a domain event |
| `PullEvents()` | Returns all registered events and clears the list |

#### `DomainEvent`

Interface that all domain events must implement.

```go
type OrderCreatedEvent struct {
    OrderID    string
    occurredAt time.Time
}

func (e OrderCreatedEvent) EventName() string     { return "order.created" }
func (e OrderCreatedEvent) OccurredAt() time.Time { return e.occurredAt }
```

#### `ValueObject`

Interface for value objects. Implementors must define equality by value, not identity.

```go
type Email struct {
    value string
}

func NewEmail(value string) (Email, error) {
    if !strings.Contains(value, "@") {
        return Email{}, errors.New("invalid email")
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
return dderr.New("USER_INACTIVE", "user is inactive", nil)

// wrapping a cause
return dderr.New("DB_ERROR", "error fetching user", err)
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
type ReportService struct {
    repo repository.Reader[string, *User]
}

// an import service only needs to write
type ImportService struct {
    repo repository.Writer[*User]
}
```

Extend with domain-specific methods:

```go
type UserRepository interface {
    repository.Repository[string, *User]
    FindByEmail(ctx context.Context, email string) (*User, error)
    FindByUsername(ctx context.Context, username string) (*User, error)
}
```

---

### `pagination`

Utilities for paginated queries.

#### `Params`

Holds pagination parameters with built-in normalization.

```go
import "github.com/aleodoni/go-ddd/pagination"

params := pagination.NewParams(page, limit)

// normalized values — page defaults to 1, limit defaults to 20 (max 100)
fmt.Println(params.Page)     // 1
fmt.Println(params.Limit)    // 20
fmt.Println(params.Offset()) // 0
```

| Method | Description |
|--------|-------------|
| `NewParams(page, limit)` | Creates normalized pagination params |
| `Normalize()` | Ensures page >= 1 and 1 <= limit <= 100 |
| `Offset()` | Returns the offset for SQL queries: `(page - 1) * limit` |

#### `PagedResult[T]`

Generic wrapper for paginated responses.

```go
result := pagination.PagedResult[*User]{
    Items: users,
    Total: 42,
    Page:  1,
    Limit: 20,
}
```

Typical use in a use case:

```go
params := pagination.NewParams(input.Page, input.Limit)

users, total, err := repo.ListUsers(ctx, input.Search, params.Page, params.Limit)
if err != nil {
    return nil, err
}

return &pagination.PagedResult[*User]{
    Items: users,
    Total: total,
    Page:  params.Page,
    Limit: params.Limit,
}, nil
```

---

## Running tests

```bash
go test ./...
```

## License

MIT