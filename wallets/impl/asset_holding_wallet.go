package impl

import (
	"errors"
	"fmt"
)

type AssetHoldingWallet struct {
	balance float32
	asset   string
	holding float32
}

func (ahw *AssetHoldingWallet) GetBalance() float32 {
	return ahw.balance
}

func (ahw *AssetHoldingWallet) Deposit(amt float32) {
	ahw.balance = ahw.balance + amt
}

func (ahw *AssetHoldingWallet) Withdraw(amt float32) error {
	newBalance := ahw.balance - amt
	if newBalance < 0 {
		return errors.New("balance cannot be less than 0")
	}
	ahw.balance = newBalance
	return nil
}

func (ahw *AssetHoldingWallet) GetHolding() float32 {
	return ahw.holding
}

func (ahw *AssetHoldingWallet) GetNetValue(price float32) (float32, error) {
	if price < 0 {
		return 0.0, fmt.Errorf("price of asset %g cannot be less than 0", price)
	}
	return ahw.balance + price*ahw.holding, nil
}

func (ahw *AssetHoldingWallet) BuyAsset(amt float32, price float32) error {
	assetValue := amt * price
	if assetValue > ahw.balance {
		return fmt.Errorf("cannot buy amount %g of %s at price %g, exceeds wallet balance %g",
			amt, ahw.asset, price, ahw.balance)
	}
	ahw.balance = ahw.balance - assetValue
	ahw.holding = ahw.holding + amt
	return nil
}

func (ahw *AssetHoldingWallet) SellAsset(amt float32, price float32) error {
	assetValue := amt * price
	if amt > ahw.holding {
		return fmt.Errorf("cannot sell amount %g of %s at price %g, exceeds wallet holding %g",
			amt, ahw.asset, price, ahw.holding)
	}
	ahw.holding = ahw.holding - amt
	ahw.balance = ahw.balance + assetValue
	return nil
}

func NewAssetHoldingWallet(balance float32, asset string, holding float32) *AssetHoldingWallet {
	return &AssetHoldingWallet{
		balance,
		asset,
		holding,
	}
}
