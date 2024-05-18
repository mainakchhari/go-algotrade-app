package impl

import (
	"go-algotrade-app/events"
	"go-algotrade-app/wallets"
)

type SimpleWallet struct {
	amt       float32
	eventChan chan events.IEvent
}

func (sw *SimpleWallet) GetEventChan() chan events.IEvent {
	return sw.eventChan
}

func (sw *SimpleWallet) GetAmount() float32 {
	return sw.amt
}

func (sw *SimpleWallet) AddTxn(txn *wallets.ITransaction) {
	sw.amt = sw.amt + (*txn).GetAmount()*float32((*txn).GetBound())
}

func NewSimpleWallet(startAmt float32, startHolding float32) wallets.IWallet {
	wallet := new(SimpleWallet)
	wallet.amt = startAmt
	return wallet
}
