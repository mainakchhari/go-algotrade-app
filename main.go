package main

import (
	"fmt"
	binance_datasource "go-algotrade-app/binance/datasource"

	"go-algotrade-app/strategy"

	binance_strategy "go-algotrade-app/binance/strategy"
	strategy_impl "go-algotrade-app/strategy/impl"
	wallets_impl "go-algotrade-app/wallets/impl"
	"sync"
)

func main() {

	stream := binance_datasource.NewBinanceRawStream("btcusdt@trade")
	var wg sync.WaitGroup
	defer wg.Wait()
	wg.Add(1)

	go stream.Process(&wg)
	smac := strategy_impl.NewSimpleMacStrategy(100, 10000)
	rds := strategy_impl.NewRandomDecisionStrategy()

	// Start with a wallet for each with 1_000_000 balance
	wallet_smac := wallets_impl.NewAssetHoldingWallet(
		float32(1000000),
		"BTC",
		float32(0),
	)
	wallet_rds := wallets_impl.NewAssetHoldingWallet(
		float32(1000000),
		"BTC",
		float32(0),
	)

	// create a Executor instance with source chan for each strat
	smaExecChan := make(chan strategy.IPriceEvent)
	smac_execution := binance_strategy.NewBinanceSourceStrategyExecutor(&smac, smaExecChan, wallet_smac)
	wg.Add(1)
	go smac_execution.Execute(&wg)

	rdsExecChan := make(chan strategy.IPriceEvent)
	rds_execution := binance_strategy.NewBinanceSourceStrategyExecutor(&rds, rdsExecChan, wallet_rds)
	wg.Add(1)
	go rds_execution.Execute(&wg)

	stored_value_smac, stored_value_rds := float32(0), float32(0)
	for {
		//read stream source & publish data event to each source chan
		message := <-stream.GetDataChan()
		smaExecChan <- message
		rdsExecChan <- message

		//read and print wallet value whenever either changes
		smacWalletNetValue, err := wallet_smac.GetNetValue(message.Price)
		if err != nil {
			panic(err)
		}
		rdsWalletNetValue, err := wallet_rds.GetNetValue(message.Price)
		if err != nil {
			panic(err)
		}
		if stored_value_rds != rdsWalletNetValue || stored_value_smac != smacWalletNetValue {
			fmt.Printf("Price Event %10.4f -> Wallet Balance -> SMAC %10.4f :: RDS %10.4f\n", message.Price, smacWalletNetValue, rdsWalletNetValue)
		}
		stored_value_rds, stored_value_smac = rdsWalletNetValue, smacWalletNetValue
	}
}
