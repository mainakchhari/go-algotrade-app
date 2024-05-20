package events

type IEvent interface {
	GetEventType() string
}

type IHasEventChan interface {
	GetEventChan() chan IEvent
}
