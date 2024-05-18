package datasource

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"go-algotrade-app/datasource"

	"github.com/gorilla/websocket"
)

type BinanceCombinedStreamMessage struct {
	Stream string           `json:"stream"`
	Data   BinanceBaseTrade `json:"data"`
}

func (m BinanceCombinedStreamMessage) Get() interface{} {
	return m
}

type BinanceCombinedStream struct {
	URI    string
	conn   *websocket.Conn
	resp   *http.Response
	Error  error
	trades chan BinanceCombinedStreamMessage
}

func (r *BinanceCombinedStream) GetDataChan() chan BinanceCombinedStreamMessage {
	return r.trades
}

func (r *BinanceCombinedStream) Process(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		_, connMessage, readErr := r.conn.ReadMessage()
		if readErr != nil {
			r.Error = readErr
			log.Fatal(r.Error)
		}

		var message BinanceCombinedStreamMessage
		err := json.Unmarshal(connMessage, &message)
		if err != nil {
			log.Fatal(err)
		}

		r.trades <- message
	}
}

func NewCombinedStream(name string) datasource.IStream[BinanceCombinedStreamMessage] {
	stream := new(BinanceCombinedStream)
	stream.URI = fmt.Sprintf("wss://fstream.binance.com/stream?streams=%s", name)
	stream.trades = make(chan BinanceCombinedStreamMessage)
	stream.conn, stream.resp, stream.Error = websocket.DefaultDialer.Dial(stream.URI, nil)
	if stream.Error != nil {
		log.Fatal(stream.Error)
	}
	return stream
}
