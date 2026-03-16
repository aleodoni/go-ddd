package domain_test

import (
	"testing"
	"time"

	"github.com/aleodoni/go-ddd/domain"
)

// --- fake event ---

type fakeEvent struct {
	name string
}

func (e fakeEvent) EventName() string     { return e.name }
func (e fakeEvent) OccurredAt() time.Time { return time.Now() }

// --- testes ---

func TestAggregateRoot_RaiseEvent_AdicionaEvento(t *testing.T) {
	agg := &domain.AggregateRoot[string]{}
	agg.RaiseEvent(fakeEvent{name: "usuario.criado"})

	events := agg.PullEvents()
	if len(events) != 1 {
		t.Errorf("esperava 1 evento, got %d", len(events))
	}
	if events[0].EventName() != "usuario.criado" {
		t.Errorf("esperava evento 'usuario.criado', got %s", events[0].EventName())
	}
}

func TestAggregateRoot_PullEvents_LimpaAposRetornar(t *testing.T) {
	agg := &domain.AggregateRoot[string]{}
	agg.RaiseEvent(fakeEvent{name: "usuario.criado"})

	agg.PullEvents()

	events := agg.PullEvents()
	if len(events) != 0 {
		t.Errorf("esperava 0 eventos após PullEvents, got %d", len(events))
	}
}

func TestAggregateRoot_RaiseEvent_MultiposEventos(t *testing.T) {
	agg := &domain.AggregateRoot[string]{}
	agg.RaiseEvent(fakeEvent{name: "usuario.criado"})
	agg.RaiseEvent(fakeEvent{name: "usuario.ativado"})

	events := agg.PullEvents()
	if len(events) != 2 {
		t.Errorf("esperava 2 eventos, got %d", len(events))
	}
}

func TestAggregateRoot_PullEvents_SemEventos(t *testing.T) {
	agg := &domain.AggregateRoot[string]{}

	events := agg.PullEvents()
	if len(events) != 0 {
		t.Errorf("esperava 0 eventos, got %d", len(events))
	}
}
