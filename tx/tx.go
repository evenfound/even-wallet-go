package tx

import "github.com/foxnut/go-hdwallet"

type TransactionInterface interface {
	Make(secret, destination, txHash string, amount int64) (TransactionInterface, error)
	GetSignedTransaction() string
	GetUnsignedTransaction() string
}

type CoinsStructure struct {
	Transaction TransactionInterface
	Type        uint32
}

var Coins = map[int]CoinsStructure{
	0: {
		Transaction: &BTCTransaction{},
		Type:        hdwallet.BTC,
	},
	2: {
		Transaction: &LTCTransaction{},
		Type:        hdwallet.LTC,
	},
}
