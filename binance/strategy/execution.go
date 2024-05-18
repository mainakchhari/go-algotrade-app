package strategy

import (
	binance_datasource "go-algotrade-app/binance/datasource"
	"go-algotrade-app/datasource"
	"go-algotrade-app/strategy"
	strategy_impl "go-algotrade-app/strategy/impl"
	"go-algotrade-app/wallets"
	wallets_impl "go-algotrade-app/wallets/impl"
	"sync"
)

type BinanceSourceStratExecution struct {
	strategy *strategy.IStrategy
	stream   datasource.IStream[binance_datasource.BinanceBaseTrade]
	wallet   *wallets.IWallet
}

func (se *BinanceSourceStratExecution) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		message := <-se.stream.GetDataChan()
		event := strategy_impl.NewEvent(message.Get().(binance_datasource.BinanceBaseTrade).Price)
		decision, price := (*se.strategy).Execute(event)
		amount := price * 0.001
		txn := wallets_impl.NewTransaction(se.wallet, amount, wallets.TxnBound(-1*decision))
		(*se.wallet).AddTxn(&txn)
	}

}

func NewStratExecution(strategy *strategy.IStrategy, stream datasource.IStream[binance_datasource.BinanceBaseTrade], wallet *wallets.IWallet) *BinanceSourceStratExecution {
	newExecution := new(BinanceSourceStratExecution)
	newExecution.strategy = strategy
	newExecution.stream = stream
	newExecution.wallet = wallet
	return newExecution
}
