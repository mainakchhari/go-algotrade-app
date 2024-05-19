package impl

type PriceEvent struct {
	price float32
}

func (e PriceEvent) GetPrice() float32 {
	return e.price
}

func NewPriceEvent(p float32) PriceEvent {
	return PriceEvent{
		price: p,
	}
}
