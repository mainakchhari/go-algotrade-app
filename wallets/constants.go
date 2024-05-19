package wallets

type TxnBound int

func (t TxnBound) Sprintf() string {
	if t == INBOUND {
		return "INBOUND"
	}
	return "OUTBOUND"
}

const (
	INBOUND  TxnBound = 1
	OUTBOUND TxnBound = -1
)
