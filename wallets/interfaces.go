package wallets

import "go-algotrade-app/events"

type IWallet interface {
	events.IHasEventChan
	GetAmount() float32
	AddTxn(*ITransaction)
}

type ITransaction interface {
	GetWallet() *IWallet
	GetAmount() float32
	GetBound() int
}
