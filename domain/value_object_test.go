package domain_test

import (
	"testing"

	"github.com/aleodoni/go-ddd/domain"
)

// --- fake value object ---

type fakeValueObject struct {
	value string
}

func (f fakeValueObject) Equals(other domain.ValueObject) bool {
	o, ok := other.(fakeValueObject)
	return ok && f.value == o.value
}

// --- testes ---

func TestValueObject_Equals_MesmosValores(t *testing.T) {
	a := fakeValueObject{value: "aleodoni"}
	b := fakeValueObject{value: "aleodoni"}

	if !a.Equals(b) {
		t.Error("esperava que a == b com o mesmo valor")
	}
}

func TestValueObject_Equals_ValoresDiferentes(t *testing.T) {
	a := fakeValueObject{value: "aleodoni"}
	b := fakeValueObject{value: "outro"}

	if a.Equals(b) {
		t.Error("esperava que a != b com valores diferentes")
	}
}

func TestValueObject_ImplementaInterface(t *testing.T) {
	var _ domain.ValueObject = fakeValueObject{}
}
