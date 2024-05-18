package impl

type Event struct {
	price float32
}

func (e *Event) GetPrice() float32 {
	return e.price
}

func NewEvent(p float32) *Event {
	return &Event{
		price: p,
	}
}
