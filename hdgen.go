package hdgen

import (
	"errors"
	"github.com/evenfound/even-wallet-go/btc"
	"github.com/evenfound/even-wallet-go/core"
	"github.com/evenfound/even-wallet-go/ltc"
	"strings"
)

const (
	errorUnsupportedCoin = "Unsupported coin"
	errorWalletConflict  = "Wallet already exists"
	errorWalletNotFound  = "Authorize error"
)

// Wallet struct is the struct of wallet which will be created
// To get/set error messages use HDWallet methods
// Also there is a opportunity to set log file path
type Wallet struct {
	core.HDWallet
}

type Account struct {
	Wallet
	AccountName string
}

// Create function creates new wallet in specified directory.
// If wallet already exists will returns conflict error
// Also if the coin type is not supported will be returned error message
func (wallet *Wallet) Create() bool {

	var w, err = wallet.identify()

	if err != nil {
		wallet.SetErrorMessage(err.Error())
		return false
	}

	if ok, _ := w.Exists(); ok {
		wallet.SetErrorMessage(errorWalletConflict)
		return false
	}

	_, err = w.Create()

	if err != nil {
		wallet.SetErrorMessage(err.Error())
		return false
	}

	return true
}

// Creating a new account
func (wallet *Account) NewAccount(name string) (accountID uint32) {

	var w, err = wallet.identify()

	if err != nil {
		wallet.SetErrorMessage(err.Error())
		return 0
	}

	accountID, err = w.NextAccount(name)

	if err != nil {
		wallet.SetErrorMessage(err.Error())
		return 0
	}

	return
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
				Coin:     wallet.Coin,
			},
		}, nil
	case "ltc":
		return ltc.Wallet{
			HDWallet: core.HDWallet{
				Name:     wallet.Name,
				Password: wallet.Password,
				Phrase:   wallet.Phrase,
				TestNet:  wallet.TestNet,
				Coin:     wallet.Coin,
			},
		}, nil
	default:
		return nil, errors.New(errorUnsupportedCoin)
	}
	return nil, nil
}
