package strategy

import (
	"go-algotrade-app/strategy"
	"go-algotrade-app/wallets"
	"sync"
)

type FixedAmountStrategyExecutor struct {
	FixedAmt float32
	strategy strategy.IStrategy
	source   chan strategy.IStrategyEvent
	wallet   wallets.IHoldAsset
}

func (se *FixedAmountStrategyExecutor) GetDataSource() chan strategy.IStrategyEvent {
	return se.source
}

func (se *FixedAmountStrategyExecutor) Execute(wg *sync.WaitGroup) {
	defer wg.Done()

	// read from source chan
	for message := range se.source {

		//execute strategy and get a decision
		decision, price := se.strategy.Execute(message)

		if decision == strategy.BUY {
			se.wallet.BuyAsset(se.FixedAmt, price)
		}

		if decision == strategy.SELL {
			se.wallet.SellAsset(se.FixedAmt, price)
		}

	}

}

func NewFixedAmountStrategyExecutor(fixedAmt float32, strat strategy.IStrategy, wallet wallets.IHoldAsset) *FixedAmountStrategyExecutor {
	source := make(chan strategy.IStrategyEvent)
	return &FixedAmountStrategyExecutor{
		fixedAmt,
		strat,
		source,
		wallet,
	}
}
