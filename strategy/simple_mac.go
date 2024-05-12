package strategy

type SimpleMAC struct {
	L200 []Event
	Chan chan Event
}

func (s *SimpleMAC) sma_l20() float64 {
	sum := float64(0)
	length := float64(len(s.L200))
	upper_bound := int(length)
	lower_bound := int(length - 20 - 1)
	if lower_bound < 0 {
		lower_bound = 0
	}
	for _, event := range s.L200[lower_bound:upper_bound] {
		sum += event.Price
	}
	return sum / length
}

func (s *SimpleMAC) Execute() {
	for {
		event := <-s.Chan
		if len(s.L200) > 0 {
			prev_event := s.L200[0]
		}
		s.L200 = append(s.L200, event)

		// if MA - L20 crosses MA - L200 up, buy
		if sma_l20(event)-sma_l200(event) > 0 && sma_l20(prev_event) < sma_l200(prev_event) {

		}

		// if MA - L20 crosses MA - L200 down, sell

		// else hold (no action)

		//
	}
}

func NewSimpleMAC(lowerLimit int, upperLimit int) *SimpleMAC {
	s := new(SimpleMAC)
	s.L200 = make([]Event, 0, 200)
	s.Chan = make(chan Event)
}
