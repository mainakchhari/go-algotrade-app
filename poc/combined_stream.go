package poc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type BinanceCombinedStreamMessage struct {
	Stream string           `json:"stream"`
	Data   BinanceBaseTrade `json:"data"`
}

type BinanceCombinedStream struct {
	URI    string
	conn   *websocket.Conn
	resp   *http.Response
	Error  error
	Trades chan BinanceCombinedStreamMessage
}

func NewCombinedStream(name string) *BinanceCombinedStream {
	stream := new(BinanceCombinedStream)
	stream.URI = fmt.Sprintf("wss://fstream.binance.com/stream?streams=%s", name)
	stream.Trades = make(chan BinanceCombinedStreamMessage)
	stream.conn, stream.resp, stream.Error = websocket.DefaultDialer.Dial(stream.URI, nil)
	if stream.Error != nil {
		log.Fatal(stream.Error)
	}
	return stream
}

func (r *BinanceCombinedStream) Process() {
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

		r.Trades <- message
	}
}
