package impl

import (
	"go-algotrade-app/strategy"
	"go-algotrade-app/wallets"
	"sync"
)

type BinanceSourceStrategyExecutor struct {
	strategy strategy.IStrategy
	source   chan strategy.IPriceEvent
	wallet   wallets.IHoldAsset
}

func (se *BinanceSourceStrategyExecutor) GetDataSource() chan strategy.IPriceEvent {
	return se.source
}

func (se *BinanceSourceStrategyExecutor) Execute(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// read from source chan
		message := <-se.source

		//execute strategy and get a decision
		decision, price := se.strategy.Execute(message)
		// fix the amount of asset to transact at 0.001
		amt := float32(0.001)

		if decision == strategy.BUY {
			se.wallet.BuyAsset(amt, price)
		}

		if decision == strategy.SELL {
			se.wallet.SellAsset(amt, price)
		}

	}

}

func NewBinanceSourceStrategyExecutor(strat strategy.IStrategy, wallet wallets.IHoldAsset) *BinanceSourceStrategyExecutor {
	source := make(chan strategy.IPriceEvent)
	return &BinanceSourceStrategyExecutor{
		strat,
		source,
		wallet,
	}
}
