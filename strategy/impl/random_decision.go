package impl

import (
	"go-algotrade-app/strategy"
	"math/rand"
	"time"
)

type RandomDecisionStrategy struct{}

func (*RandomDecisionStrategy) Execute(event strategy.IPriceEvent) (strategy.Decision, float32) {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	decision := strategy.Decision(rand.Intn(3) - 1)
	return decision, event.GetPrice()
}

func NewRandomDecisionStrategy() RandomDecisionStrategy {
	return RandomDecisionStrategy{}
}
