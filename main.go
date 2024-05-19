package main

import (
	"fmt"
	binance_datasource "go-algotrade-app/binance/datasource"
	"os"

	binance_strategy "go-algotrade-app/binance/strategy"
	strategy_impl "go-algotrade-app/strategy/impl"
	wallets_impl "go-algotrade-app/wallets/impl"
	"sync"

	"github.com/jedib0t/go-pretty/v6/table"
)

func formatFloat(f float32) string {
	return fmt.Sprintf("%10.4f", f)
}

func initTableWriter() *table.Writer {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Strategy", "Balance", "Holding", "Net Val"})
	return &t
}

func renderTable(t *table.Writer, tr []table.Row) {
	(*t).ResetRows()
	(*t).AppendRows(tr)
	// hack to clear/refresh screen
	fmt.Print("\033[H\033[2J")
	(*t).Render()
}

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

	// create a Executor instance for each strat
	smac_execution := binance_strategy.NewBinanceSourceStrategyExecutor(&smac, wallet_smac)
	wg.Add(1)
	go smac_execution.Execute(&wg)

	rds_execution := binance_strategy.NewBinanceSourceStrategyExecutor(&rds, wallet_rds)
	wg.Add(1)
	go rds_execution.Execute(&wg)

	// init table for display
	tw := initTableWriter()

	for {
		//read stream source & publish data event to each source chan
		message := <-stream.GetDataSink()
		smac_execution.GetDataSource() <- message
		rds_execution.GetDataSource() <- message

		//read and print wallet value whenever either changes
		smacWalletNetValue, _ := wallet_smac.GetNetValue(message.Price)
		smacWalletBalance := wallet_smac.GetBalance()
		smacWalletHolding := wallet_smac.GetHolding()
		rdsWalletNetValue, _ := wallet_rds.GetNetValue(message.Price)
		rdsWalletBalance := wallet_rds.GetBalance()
		rdsWalletHolding := wallet_rds.GetHolding()
		renderTable(tw, []table.Row{
			{"RDS", formatFloat(rdsWalletBalance), formatFloat(rdsWalletHolding), formatFloat(rdsWalletNetValue)},
			{"SMAC", formatFloat(smacWalletBalance), formatFloat(smacWalletHolding), formatFloat(smacWalletNetValue)},
		})
	}
}
