package datasource

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

const BINANCE_WEBSOCKET_RAW_BASE = "wss://fstream.binance.com/ws/%s"

// BinanceRawStream represents a raw stream for Binance trades.
type BinanceRawStream struct {
	URI    string
	conn   *websocket.Conn
	resp   *http.Response
	trades chan BinanceBaseTrade
}

func (r *BinanceRawStream) GetDataSink() chan BinanceBaseTrade {
	return r.trades
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

// Process reads messages from the BinanceRawStream connection and processes them.
// It unmarshals the received messages and sends them to the Trades channel.
// If there is an error while reading or unmarshaling the messages, it sets the Error field and logs the error.
func (r *BinanceRawStream) Process(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		_, connMessage, readErr := r.conn.ReadMessage()
		if readErr != nil {
			close(r.trades)
			log.Fatal(readErr)
		}

		messages, err := r.unmarshalArray(connMessage)
		if err != nil {
			message, err := r.unmarshalObject(connMessage)
			if err != nil {
				close(r.trades)
				log.Fatal(err)
			}
			r.trades <- *message
			continue
		}
		for _, message := range *messages {
			r.trades <- message
		}
	}
}

// NewBinanceRawStream creates a new instance of BinanceRawStream with the specified name.
// It establishes a WebSocket connection to the Binance raw stream API and returns the stream object.
func NewBinanceRawStream(name string) *BinanceRawStream {
	streamURI := fmt.Sprintf(BINANCE_WEBSOCKET_RAW_BASE, name)
	conn, resp, err := websocket.DefaultDialer.Dial(streamURI, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &BinanceRawStream{
		streamURI,
		conn,
		resp,
		make(chan BinanceBaseTrade),
	}
}
