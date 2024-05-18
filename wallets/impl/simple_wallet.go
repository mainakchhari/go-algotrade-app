package impl

import "go-algotrade-app/wallets"

type SimpleWallet struct {
	amt     float32
	holding float32
	txns    []wallets.ITxn
}

func (sw *SimpleWallet) GetAmount() float32 {
	return sw.amt
}

func (sw *SimpleWallet) GetHolding() float32 {
	return sw.holding
}

func (sw *SimpleWallet) GetTxns() []wallets.ITxn {
	return sw.txns
}

func (sw *SimpleWallet) AddTxn(txn *wallets.ITxn) []wallets.ITxn {
	sw.txns = append(sw.txns, *txn)
	sw.amt = sw.amt + (*txn).GetAmount()*float32((*txn).GetBound())
	return sw.txns
}

func NewSimpleWallet(startAmt float32, startHolding float32) wallets.IWallet {
	wallet := new(SimpleWallet)
	wallet.amt = startAmt
	wallet.holding = startHolding
	wallet.txns = make([]wallets.ITxn, 0)
	return wallet
}

type Transaction struct {
	wallet *wallets.IWallet
	amount float32
	bound  wallets.TxnBound
}

func (t *Transaction) GetWallet() *wallets.IWallet {
	return t.wallet
}

func (t *Transaction) GetAmount() float32 {
	return t.amount
}

func (t *Transaction) GetBound() int {
	return int(t.bound)
}

func NewTransaction(wallet *wallets.IWallet, amount float32, bound wallets.TxnBound) wallets.ITxn {
	txn := new(Transaction)
	txn.wallet = wallet
	txn.bound = bound
	txn.amount = amount
	return txn
}
