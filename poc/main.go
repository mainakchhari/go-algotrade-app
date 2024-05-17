package poc

import (
	"fmt"
	"log"
	"sync"
)

// URI := "wss://fstream.binance.com/stream?streams=bnbusdt@aggTrade/btcusdt@markPrice"
// URI := "wss://fstream.binance.com/ws/!markPrice@arr@1s"
// URI := "wss://fstream.binance.com/stream?streams=!markPrice@arr@1s"

func streamPoc(wg *sync.WaitGroup) {
	defer wg.Done()

	stream := NewRawStream("btcusdt@trade")

	go stream.Process()

	for {
		trade := <-stream.Trades
		if stream.Error != nil {
			log.Fatal(stream.Error)
		}
		fmt.Printf("Time %s %+v\n", trade.DisplayTime(), trade)
	}
}

func combinedStreamPoc(wg *sync.WaitGroup) {
	defer wg.Done()
	stream := NewCombinedStream("bnbusdt@aggTrade/btcusdt@markPrice")

	go stream.Process()

	for {
		message := <-stream.Trades
		if stream.Error != nil {
			log.Fatal(stream.Error)
		}
		fmt.Printf("Time %s %+v\n", message.Data.DisplayTime(), message)
	}
}

func MainPoc() {
	var wg sync.WaitGroup
	wg.Add(2)

	go streamPoc(&wg)
	go combinedStreamPoc(&wg)

	wg.Wait()
}

func main() {
	MainPoc()
}
