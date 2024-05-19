package wallets

type IWallet interface {
	GetBalance() float32
	Deposit(float32)
	Withdraw(float32) error
}

type IHoldAsset interface {
	GetHolding() float32
	GetNetValue(price float32) (float32, error)
	BuyAsset(amt float32, price float32) error
	SellAsset(amt float32, price float32) error
}
