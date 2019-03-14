package btc

import (
	"github.com/btcsuite/btcwallet/netparams"
	btcdWallet "github.com/btcsuite/btcwallet/wallet"
	"github.com/joho/godotenv"
)

const coinName = "btc"

func init() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

}

// Getting wallet loader
func (wallet *Wallet) GetLoader() *btcdWallet.Loader {

	var network = wallet.GetNetwork()

	var path = wallet.Path()

	return btcdWallet.NewLoader(network.Params, path, 0255)
}

// Getting btcd-network by network type
func (wallet *Wallet) GetNetwork() netparams.Params {

	if wallet.TestNet {
		return netparams.TestNet3Params
	}

	return netparams.MainNetParams
}
