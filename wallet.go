package hdgen

import (
	"encoding/hex"
	"errors"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/evenfound/even-wallet-go/tx"
	"github.com/foxnut/go-hdwallet"
)

const (
	UnAuthorized = "Unauthorized"
)

type HDWallet struct {
	errorMessage error

	masterWallet *hdwallet.Key
	baseWallet   *hdwallet.Wallet

	testNet bool
}

// setError function sets an error
func (wallet *HDWallet) setError(err error) {
	wallet.errorMessage = err
}

// GetErrorMessage parse an error message from error object and returns it to user
// If there is an't  error than user will be receive empty string
func (wallet *HDWallet) GetErrorMessage() string {
	if wallet.errorMessage != nil {
		return wallet.errorMessage.Error()
	}
	return ""
}

// Creating a new seed hash from mnemonic phrase and password
// If an error caused that user will be receive an empty string and error message
// using GetErrorMessage, otherwise will be receive seed hash
func (wallet *HDWallet) Create(mnemonic, password string) string {

	var seed, err = hdwallet.NewSeed(mnemonic, password, "")

	if err != nil {
		wallet.setError(err)
	} else {
		return hex.EncodeToString(seed)
	}

	return ""
}

// Authorize function authorizes given mnemonic phrase by validation it via BIP39 standards
// If an error caused user will be receive an error message using GetErrorMessage function, otherwise
// in the object will be available master wallet
func (wallet *HDWallet) Authorize(seed string) bool {
	var hash, err = hex.DecodeString(seed)

	if err != nil {
		wallet.setError(err)
		return false
	}

	master, err := hdwallet.NewKey(hdwallet.Seed(hash))

	if err != nil {
		wallet.setError(err)
		return false
	}

	wallet.masterWallet = master

	return true

}

// NewAddress function generates a new address based on coin / account / change / address_index
// If wallet is not  already authorized  that user will be get an error
// otherwise user will be receive an address
func (wallet *HDWallet) NewAddress(coin, account, change, address int) (addr string) {

	coinData, ok := tx.Coins[coin]
	if !ok {
		wallet.setError(errors.New(UnsupportedCoinType))
		return ""
	}

	// Converting int to uint32
	addressUint32 := uint32(address)
	changeUint32 := uint32(change)
	accountUint32 := uint32(account)

	hdwallet.CoinType(hdwallet.BTC)

	if wallet.masterWallet == nil {
		wallet.setError(errors.New(UnAuthorized))
	} else {
		var params = chaincfg.MainNetParams

		if wallet.testNet {
			params = chaincfg.TestNet3Params
		}

		var hdWallet, err = wallet.masterWallet.GetWallet(
			hdwallet.Params(&params),
			hdwallet.CoinType(coinData.Type),
			hdwallet.Account(accountUint32),
			hdwallet.Change(changeUint32),
			hdwallet.AddressIndex(addressUint32),
		)

		if err != nil {
			wallet.setError(err)
		} else {
			addr, err = hdWallet.GetAddress()
			if err != nil {
				addr = ""
				wallet.setError(err)
				//
			}
		}
	}
	return
}

// WIF function returns WIF (Wallet import format) to sign transactions
func (wallet *HDWallet) WIF(coin, account, change, address int) string {
	addressUint32 := uint32(address)
	changeUint32 := uint32(change)
	accountUint32 := uint32(account)
	var coinType uint32
	var coinData, ok = tx.Coins[coin]

	if ok {
		coinType = coinData.Type
		if wallet.masterWallet == nil {
			wallet.setError(errors.New("Unauthorized"))
		} else {
			var key, err = wallet.masterWallet.GetChildKey(
				hdwallet.CoinType(coinType),
				hdwallet.Account(accountUint32),
				hdwallet.Change(changeUint32),
				hdwallet.AddressIndex(addressUint32),
			)
			wif, err := key.PrivateWIF(false)
			if err != nil {
				wallet.setError(err)
			} else {
				return wif
			}
		}
	} else {
		wallet.setError(errors.New(UnsupportedCoinType))
	}

	return ""
}

func (wallet *HDWallet) IsTestNet(testNet bool) {
	wallet.testNet = testNet
}
