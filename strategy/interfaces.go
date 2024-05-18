package strategy

import "go-algotrade-app/events"

// A IStrategy is a set of rules that determine when to buy, sell, or hold an asset.
// It is a struct that contains the rules and the data that the rules operate on.
// The rules are implemented as methods on the struct.
// The data is stored as fields on the struct.
// The struct also contains a method that executes the rules.
// The struct is created by a constructor function.
// The struct is passed to a function that executes the rules.
type IStrategy interface {
	Execute(event IPriceEvent) (Decision, float32)
}

type IPriceEvent interface {
	events.IEvent
	GetPrice() float32
}
