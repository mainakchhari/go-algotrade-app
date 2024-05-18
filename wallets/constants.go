package wallets

type TxnBound int

const (
	INBOUND  TxnBound = 1
	OUTBOUND TxnBound = -1
)
