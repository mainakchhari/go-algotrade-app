package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type BinanceBaseTrade struct {
	EventType            string  `json:"e"`
	EventTime            int64   `json:"E"`
	AggTradeId           int64   `json:"a"` // aggTrade only
	Symbol               string  `json:"s"`
	Price                float32 `json:"p,string"`
	IndexPrice           float32 `json:"i,string"` // markPriceUpdate only
	EstimatedSettlePrice float32 `json:"P,string"` // markPriceUpdate only; Estimated Settle Price, only useful in the last hour before the settlement starts
	Quantity             float32 `json:"q,string"` // aggTrade only
	FirstTradeId         int64   `json:"f"`        // aggTrade only
	LastTradeId          int64   `json:"l"`        // aggTrade only
	FundingRate          float32 `json:"r,string"` // markPriceUpdate only
	TradeTime            int64   `json:"T"`        // for markPriceUpdate, Next funding time
	IsMarketMaker        bool    `json:"m"`        // aggTrade only
}

func (t *BinanceBaseTrade) DisplayTime() *time.Time {
	if t.TradeTime > 0 {
		displayTime := time.Unix(t.TradeTime/1000, (t.TradeTime%1000)*1000000)
		return &displayTime
	} else {
		displayTime := time.Unix(t.TradeTime/1000, (t.TradeTime%1000)*1000000)
		return &displayTime
	}
}

type BinanceCombinedStreamMessage struct {
	Stream string             `json:"stream"`
	Data   []BinanceBaseTrade `json:"data"`
}

type BinanceRawStream struct {
	URI    string
	conn   *websocket.Conn
	resp   *http.Response
	error  error
	Trades chan BinanceBaseTrade
}

func NewRawStream(name string) *BinanceRawStream {
	stream := new(BinanceRawStream)
	stream.URI = fmt.Sprintf("wss://fstream.binance.com/ws/%s", name)
	stream.Trades = make(chan BinanceBaseTrade)
	stream.conn, stream.resp, stream.error = websocket.DefaultDialer.Dial(stream.URI, nil)
	if stream.error != nil {
		log.Fatal(stream.error)
	}
	return stream
}

func (r *BinanceRawStream) unmarshalArray(byteArray []byte) (*[]BinanceBaseTrade, error) {
	var messages []BinanceBaseTrade
	err := json.Unmarshal(byteArray, &messages)
	if err != nil {
		return &messages, fmt.Errorf("%v:\n%v", err, string(byteArray))
	}
	return &messages, nil
}

func (r *BinanceRawStream) unmarshalObject(byteArray []byte) (*BinanceBaseTrade, error) {
	var message BinanceBaseTrade
	err := json.Unmarshal(byteArray, &message)
	if err != nil {
		return &message, fmt.Errorf("%v:\n%v", err, string(byteArray))
	}
	return &message, nil
}

func (r *BinanceRawStream) Process() {
	for {
		_, connMessage, readErr := r.conn.ReadMessage()
		if readErr != nil {
			r.error = readErr
			log.Fatal(r.error)
		}

		messages, err := r.unmarshalArray(connMessage)
		if err != nil {
			message, err := r.unmarshalObject(connMessage)
			if err != nil {
				r.error = err
				log.Fatal(r.error)
			}
			r.Trades <- *message
			continue
		}
		for _, message := range *messages {
			r.Trades <- message
		}
	}
}

func main() {
	// URI := "wss://fstream.binance.com/stream?streams=bnbusdt@aggTrade/btcusdt@markPrice"
	// URI := "wss://fstream.binance.com/ws/!markPrice@arr@1s"
	// URI := "wss://fstream.binance.com/stream?streams=!markPrice@arr@1s"

	stream := NewRawStream("btcusdt@trade")

	go stream.Process()

	for {
		trade := <-stream.Trades
		if stream.error != nil {
			log.Fatal(stream.error)
		}
		fmt.Printf("Time %s %+v\n", trade.DisplayTime(), trade)
	}
}
