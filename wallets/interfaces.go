package wallets

type IWallet interface {
	GetAmount() float32
	GetHolding() float32
	GetTxns() []ITxn
	AddTxn(*ITxn) []ITxn
}

type ITxn interface {
	GetWallet() *IWallet
	GetAmount() float32
	GetBound() int
}
