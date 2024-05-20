package datasource

import "time"

type BinanceBaseTrade struct {
	EventType            string  `json:"e"`
	EventTime            int64   `json:"E"`
	AggTradeId           int64   `json:"a"` // aggTrade only
	Symbol               string  `json:"s"`
	Price                float32 `json:"p,string"`
	TradeId              int64   `json:"t"`        // trade only
	IndexPrice           float32 `json:"i,string"` // markPriceUpdate only
	EstimatedSettlePrice float32 `json:"P,string"` // markPriceUpdate only; Estimated Settle Price, only useful in the last hour before the settlement starts
	Quantity             float32 `json:"q,string"` // aggTrade only
	FirstTradeId         int64   `json:"f"`        // aggTrade only
	LastTradeId          int64   `json:"l"`        // aggTrade only
	FundingRate          float32 `json:"r,string"` // markPriceUpdate only
	TradeTime            int64   `json:"T"`        // for markPriceUpdate, Next funding time
	IsMarketMaker        bool    `json:"m"`        // aggTrade only
}

func (t BinanceBaseTrade) GetEventType() string {
	return t.EventType
}

func (t BinanceBaseTrade) Get() interface{} {
	return t
}

func (t BinanceBaseTrade) GetPrice() float32 {
	return t.Price
}

func (t BinanceBaseTrade) GetQuantity() float32 {
	return t.Quantity
}

func (t *BinanceBaseTrade) DisplayTime() *time.Time {
	if t.TradeTime > 0 {
		displayTime := time.Unix(t.TradeTime/1000, (t.TradeTime%1000)*1000000)
		return &displayTime
	} else {
		displayTime := time.Unix(t.EventTime/1000, (t.EventTime%1000)*1000000)
		return &displayTime
	}
}
