package domain

type AggregateRoot[ID any] struct {
	Entity[ID]
	events []DomainEvent
}

func (a *AggregateRoot[ID]) RaiseEvent(e DomainEvent) {
	a.events = append(a.events, e)
}

func (a *AggregateRoot[ID]) PullEvents() []DomainEvent {
	evts := a.events
	a.events = nil
	return evts
}

func NewAggregateRoot[ID any](id ID) AggregateRoot[ID] {
	return AggregateRoot[ID]{
		Entity: Entity[ID]{ID: id},
	}
}
