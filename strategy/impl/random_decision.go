package impl

import (
	"go-algotrade-app/strategy"
	"math/rand"
	"time"
)

type RandomDecisionStrategy struct{}

func (*RandomDecisionStrategy) Execute(event strategy.IEvent) (strategy.Decision, float32) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	decision := strategy.Decision(rand.Intn(3) - 1)
	return decision, event.GetPrice()
}

func NewRandomDecisionStrategy() strategy.IStrategy {
	return new(RandomDecisionStrategy)
}
