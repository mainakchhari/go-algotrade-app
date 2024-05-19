package datasource

type BinanceCombinedStreamMessage struct {
	Stream string           `json:"stream"`
	Data   BinanceBaseTrade `json:"data"`
}

func (m BinanceCombinedStreamMessage) Get() interface{} {
	return m
}
