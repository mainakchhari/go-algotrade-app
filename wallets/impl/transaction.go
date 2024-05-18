package impl

import "go-algotrade-app/wallets"

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

func NewTransaction(wallet *wallets.IWallet, amount float32, bound wallets.TxnBound) wallets.ITransaction {
	txn := new(Transaction)
	txn.wallet = wallet
	txn.bound = bound
	txn.amount = amount
	return txn
}
