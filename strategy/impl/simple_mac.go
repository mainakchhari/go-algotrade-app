package impl

import (
	"container/ring"
	"fmt"
	"go-algotrade-app/strategy"
)

type SimpleMacStrategy struct {
	shortLength int
	longRing    *ring.Ring
	shortAvg    float32
	longAvg     float32
}

func (s *SimpleMacStrategy) Execute(event strategy.IStrategyEvent) (strategy.Decision, float32) {
	decision := strategy.HOLD

	switch eventType := event.GetEventType(); eventType {
	case "trade":
		s.longRing.Value = event.GetPrice()
	case "aggTrade":
		s.longRing.Value = [2]float32{event.GetPrice(), event.GetQuantity()}
	}
	s.longRing = s.longRing.Next()

	newShortAvg, newLongAvg := float32(0), float32(0)
	switch eventType := event.GetEventType(); eventType {
	case "trade":
		newShortAvg, newLongAvg = s.calcTradeAvgs()
	case "aggTrade":
		newShortAvg, newLongAvg = s.calcAggTradeAvgs()
	default:
		panic(fmt.Sprintf("cannot handle eventtype %s", eventType))
	}

	if newShortAvg > newLongAvg && s.shortAvg < s.longAvg {
		decision = strategy.BUY
	}
	if newShortAvg < newLongAvg && s.shortAvg > s.longAvg {
		decision = strategy.SELL
	}
	s.shortAvg = newShortAvg
	s.longAvg = newLongAvg
	return decision, event.GetPrice()
}

func (s *SimpleMacStrategy) calcTradeAvgs() (float32, float32) {
	shortSum, longSum := float32(0), float32(0)
	for i := 0; i < s.longRing.Len(); i++ {
		if val, ok := s.longRing.Value.(float32); ok {
			if i < s.shortLength {
				shortSum += val
			}
			longSum += val
		}
		s.longRing = s.longRing.Next()
	}
	newShortAvg := shortSum / float32(s.shortLength)
	newLongAvg := longSum / float32(s.longRing.Len())
	return newShortAvg, newLongAvg
}

func (s *SimpleMacStrategy) calcAggTradeAvgs() (float32, float32) {
	shortSum, longSum := float32(0), float32(0)
	shortQSum, longQSum := float32(0), float32(0)
	for i := 0; i < s.longRing.Len(); i++ {
		if val, ok := s.longRing.Value.([2]float32); ok {
			if i < s.shortLength {
				shortSum += val[0] * val[1]
				shortQSum += val[1]
			}
			longSum += val[0] * val[1]
			longQSum += val[1]
		}
		s.longRing = s.longRing.Next()
	}
	if shortQSum == 0 {
		shortQSum = float32(s.shortLength)
	}
	if longQSum == 0 {
		longQSum = float32(s.longRing.Len())
	}
	newShortAvg := shortSum / shortQSum
	newLongAvg := longSum / longQSum
	return newShortAvg, newLongAvg
}

func NewSimpleMacStrategy(shortPeriods int, longPeriods int) SimpleMacStrategy {
	strat := new(SimpleMacStrategy)
	strat.longRing = ring.New(longPeriods)
	strat.shortLength = shortPeriods
	return *strat
}
