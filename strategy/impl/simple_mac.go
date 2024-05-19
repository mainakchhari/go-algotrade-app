package impl

import (
	"container/ring"
	"go-algotrade-app/strategy"
)

type SimpleMacStrategy struct {
	shortLength int
	longRing    *ring.Ring
	shortAvg    float32
	longAvg     float32
}

func (s *SimpleMacStrategy) Execute(event strategy.IPriceEvent) (strategy.Decision, float32) {
	decision := strategy.HOLD

	s.longRing.Value = event.GetPrice()
	s.longRing = s.longRing.Next()

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

func NewSimpleMacStrategy(shortPeriods int, longPeriods int) SimpleMacStrategy {
	strat := new(SimpleMacStrategy)
	strat.longRing = ring.New(longPeriods)
	strat.shortLength = shortPeriods
	return *strat
}
