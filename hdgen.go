package hdgen

import (
	"errors"
	"hdgen/btc"
	"hdgen/core"
	"hdgen/ltc"
	"strings"
)

const (
	errorUnsupportedCoin = "Unsupported coin"
	errorWalletConflict  = "Wallet already exists"
)

type Wallet struct {
	core.HDWallet
	Coin string
}




func (wallet *Wallet) Create() bool {

	var w, err = wallet.identify()

	if err != nil {
		return false
	}

	if ok, _ := w.Exists(); ok {
		return false
	}

	_, err = w.Create()

	if err != nil {
		return false
	}

	return true
}

// Identifying coin type to get a specific interface
func (wallet *Wallet) identify() (core.WalletInterface, error) {

	switch strings.ToLower(wallet.Coin) {
	case "btc":
		return btc.Wallet{
			HDWallet: core.HDWallet{
				Name:     wallet.Name,
				Password: wallet.Password,
				Phrase:   wallet.Phrase,
				TestNet:  wallet.TestNet,
			},
		}, nil
	case "ltc":
		return ltc.Wallet{
			HDWallet: core.HDWallet{
				Name:     wallet.Name,
				Password: wallet.Password,
				Phrase:   wallet.Phrase,
				TestNet:  wallet.TestNet,
			},
		}, nil
	default:
		return nil, errors.New(errorUnsupportedCoin)

	}
	return nil, nil
}

func main() {}
