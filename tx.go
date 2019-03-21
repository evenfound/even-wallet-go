package hdgen

import (
	"github.com/evenfound/even-wallet-go/tx"
)

type Transaction struct {
	errorMessage string

	unsignedTx string
	signedTx   string
}

// Make function makes as well as  signs transaction in offline mode
// If the transaction successfully signed than the user will receive an signed and unsigned
// transaction hashes otherwise an error message
func (tr *Transaction) Make(coin int, secret, destination, txHash string, amount int64) {
	if data, ok := tx.Coins[coin]; ok {

		var txs, err = data.Make(secret, destination, txHash, amount)

		if err == nil {
			tr.signedTx = txs.GetSignedTransaction()
			tr.unsignedTx = txs.GetUnsignedTransaction()
		} else {
			tr.errorMessage = err.Error()
		}
	} else {
		tr.errorMessage = "Unsupported coin type"
	}
}

// GetUnsignedTransaction return unsigned transaction hash
func (tr *Transaction) GetUnsignedTransaction() string {
	return tr.unsignedTx
}

// GetSignedTransaction return signed transaction hash
func (tr *Transaction) GetSignedTransaction() string {
	return tr.signedTx
}

// GetErrorMessage returns an error message
func (tr Transaction) GetErrorMessage() string {
	return tr.errorMessage
}
