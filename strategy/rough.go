package strategy

import (
	"math/rand"
	"sync"
	"time"
)

type Decision int

const (
	SELL Decision = iota - 1
	HOLD
	BUY
)

// A Strategy is a set of rules that determine when to buy, sell, or hold an asset.
// It is a struct that contains the rules and the data that the rules operate on.
// The rules are implemented as methods on the struct.
// The data is stored as fields on the struct.
// The struct also contains a method that executes the rules.
// The struct is created by a constructor function.
// The struct is passed to a function that executes the rules.
type Strategy interface {
	Execute(wg *sync.WaitGroup, event IEvent) Decision
}

type RandomDecisionStrategy struct{}

func (*RandomDecisionStrategy) Execute(wg *sync.WaitGroup, event IEvent) Decision {
	defer wg.Done()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return Decision(rand.Intn(3) - 1)
}

func NewRandomDecisionStrategy() Strategy {
	return new(RandomDecisionStrategy)
}

type SimpleMacStrategy struct{}

func (*SimpleMacStrategy) Execute(wg *sync.WaitGroup, event IEvent) Decision {
	defer wg.Done()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return Decision(rand.Intn(3) - 1)
}

func NewSimpleMacStrategy() Strategy {
	return new(RandomDecisionStrategy)
}
