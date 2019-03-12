package ltc

import (
	"fmt"
	"github.com/btcsuite/btcwallet/netparams"
	btcdWallet "github.com/btcsuite/btcwallet/wallet"
	"github.com/joho/godotenv"
	"hdgen/core"
	"os"
	"path/filepath"
)

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

	fmt.Println(path)

	return btcdWallet.NewLoader(network.Params, path, 0255)
}

// Getting btcd-network by network type
func (wallet *Wallet) GetNetwork() netparams.Params {

	if wallet.TestNet {
		return netparams.TestNet3Params
	}

	return netparams.MainNetParams
}

// Getting wallet path
func (wallet Wallet) Path() string {

	var dbPath = os.Getenv(core.DBPath)

	var walletPath = os.Getenv(core.WalletPathKey)

	return filepath.Join(dbPath, walletPath, wallet.Name)
}
