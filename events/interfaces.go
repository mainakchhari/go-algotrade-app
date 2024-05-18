package events

type IEvent interface {
}

type IHasEventChan interface {
	GetEventChan() chan IEvent
}
