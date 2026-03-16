// Package errors provides standard domain errors for DDD applications.
package errors_test

import (
	"testing"

	dderr "github.com/aleodoni/go-ddd/errors"
)

func TestDomainError_Error_RetornaMensagem(t *testing.T) {
	err := dderr.New("USER_INACTIVE", "usuário inativo", nil)

	if err.Error() != "usuário inativo" {
		t.Errorf("esperava 'usuário inativo', got %s", err.Error())
	}
}

func TestDomainError_Unwrap_RetornaCausa(t *testing.T) {
	causa := dderr.ErrNotFound
	err := dderr.New("USER_NOT_FOUND", "usuário não encontrado", causa)

	if err.Unwrap() != causa {
		t.Errorf("esperava causa %v, got %v", causa, err.Unwrap())
	}
}

func TestDomainError_Unwrap_SemCausa(t *testing.T) {
	err := dderr.New("USER_INACTIVE", "usuário inativo", nil)

	if err.Unwrap() != nil {
		t.Errorf("esperava nil, got %v", err.Unwrap())
	}
}

func TestDomainError_Code(t *testing.T) {
	err := dderr.New("USER_INACTIVE", "usuário inativo", nil)

	if err.Code != "USER_INACTIVE" {
		t.Errorf("esperava 'USER_INACTIVE', got %s", err.Code)
	}
}

func TestErrosSentinela_NaoNil(t *testing.T) {
	erros := []error{
		dderr.ErrNotFound,
		dderr.ErrAlreadyExists,
		dderr.ErrForbidden,
		dderr.ErrUnauthorized,
		dderr.ErrInvalidInput,
		dderr.ErrConflict,
		dderr.ErrInternal,
	}

	for _, err := range erros {
		if err == nil {
			t.Errorf("esperava erro não nil, got nil para %v", err)
		}
	}
}
