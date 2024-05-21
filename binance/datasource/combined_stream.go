package datasource

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

const BINANCE_WEBSOCKET_COMBINED_BASE = "wss://fstream.binance.com/stream?streams=%s"

type BinanceCombinedStream struct {
	URI    string
	conn   *websocket.Conn
	resp   *http.Response
	trades chan BinanceCombinedStreamMessage
}

func (r *BinanceCombinedStream) GetDataSink() chan BinanceCombinedStreamMessage {
	return r.trades
}

func (r *BinanceCombinedStream) Process(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		_, connMessage, readErr := r.conn.ReadMessage()
		if readErr != nil {
			close(r.trades)
			log.Fatal(readErr)
		}

		var message BinanceCombinedStreamMessage
		err := json.Unmarshal(connMessage, &message)
		if err != nil {
			close(r.trades)
			log.Fatal(err)
		}

		r.trades <- message
	}
}

func NewBinanceCombinedStream(name string) *BinanceCombinedStream {
	streamURI := fmt.Sprintf(BINANCE_WEBSOCKET_COMBINED_BASE, name)
	conn, resp, err := websocket.DefaultDialer.Dial(streamURI, nil)
	if err != nil {
		log.Fatal(err)
	}
	return &BinanceCombinedStream{
		streamURI,
		conn,
		resp,
		make(chan BinanceCombinedStreamMessage),
	}
}
