package impl

import "errors"

type SimpleWallet struct {
	balance float32
}

func (sw *SimpleWallet) GetBalance() float32 {
	return sw.balance
}

func (sw *SimpleWallet) Deposit(amt float32) {
	sw.balance = sw.balance + amt
}

func (sw *SimpleWallet) Withdraw(amt float32) error {
	newBalance := sw.balance - amt
	if newBalance < 0 {
		return errors.New("balance cannot be less than 0")
	}
	sw.balance = newBalance
	return nil
}

func NewSimpleWallet(balance float32) SimpleWallet {
	return SimpleWallet{
		balance,
	}
}
