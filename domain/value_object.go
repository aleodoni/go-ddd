package domain

type ValueObject interface {
	Equals(other ValueObject) bool
}
