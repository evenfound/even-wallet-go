package tx

type TransactionInterface interface {
	Make(secret, destination, txHash string, amount int64) (TransactionInterface, error)
	GetSignedTransaction() string
	GetUnsignedTransaction() string
}

var Coins = map[int]TransactionInterface{
	0: &BTCTransaction{},
	2: &LTCTransaction{},
}
