package main

import (
	"fmt"
	binance_datasource "go-algotrade-app/datasource/impl/binance"

	strategy_impl "go-algotrade-app/strategy/impl"
	binance_strategy "go-algotrade-app/strategy/impl/binance"
	wallets_impl "go-algotrade-app/wallets/impl"
	"sync"
)

func main() {

	stream := binance_datasource.NewRawStream("btcusdt@trade")
	var wg sync.WaitGroup
	wg.Add(1)

	go stream.Process(&wg)
	smac := strategy_impl.NewSimpleMacStrategy(10, 100)
	rds := strategy_impl.NewRandomDecisionStrategy()

	// Start with 1000000 capital for each
	wallet_smac := wallets_impl.NewSimpleWallet(
		float32(1000000),
		float32(0),
	)
	wallet_rds := wallets_impl.NewSimpleWallet(
		float32(1000000),
		float32(0),
	)
	smac_execution := binance_strategy.NewStratExecution(&smac, stream, &wallet_smac)
	wg.Add(1)
	go smac_execution.Execute(&wg)
	rds_execution := binance_strategy.NewStratExecution(&rds, stream, &wallet_rds)
	wg.Add(1)
	go rds_execution.Execute(&wg)

	for {
		message := <-stream.GetDataChan()
		fmt.Printf("Price Event %10.3f -> Wallet Balance -> SMAC %10.3f :: RDS %10.3f\n", message.Price, wallet_smac.GetAmount(), wallet_rds.GetAmount())
	}

}
