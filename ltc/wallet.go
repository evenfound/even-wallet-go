package ltc

import (
	"github.com/ltcsuite/ltcwallet/waddrmgr"
	ltcdWallet "github.com/ltcsuite/ltcwallet/wallet"
	_ "github.com/ltcsuite/ltcwallet/walletdb/bdb"
	"hdgen/core"
	"os"
	"path/filepath"
	"time"
)

type Wallet struct {
	core.HDWallet
}

var (
	_ core.WalletInterface = (*Wallet)(nil)
	_ core.BTCDInterface   = (*ltcdWallet.Wallet)(nil)
	_ core.BTCDInterface   = ltcdWallet.Wallet{}
)

func (wallet Wallet) Create() (core.BTCDInterface, error) {

	var loader = wallet.GetLoader()

	var newWallet, err = loader.CreateNewWallet([]byte(wallet.Password), []byte(wallet.Phrase), nil, time.Now())

	defer newWallet.Manager.Close()

	if err != nil {
		return nil, err
	}

	var privateDataPath = filepath.Join(wallet.Path(), os.Getenv(core.PrivateDataFName))

	f, err := os.Create(privateDataPath)

	defer f.Close()

	return newWallet, nil
}

func (wallet Wallet) Authorize() (core.BTCDInterface, error) {
	var loader = wallet.GetLoader()

	return loader.OpenExistingWallet([]byte(wallet.Password), false)
}

func (wallet Wallet) Exists() (bool, error) {
	return wallet.GetLoader().WalletExists()
}

// Creating a new account
func (wallet Wallet) NextAccount(name string) (uint32, error) {
	var w, err = wallet.Authorize()

	// Converting to real type to get a real interface
	coreWallet, err := w.(ltcdWallet.Wallet)

	if err != nil {
		return 0, nil
	}

	var keyScope = getKeyScope()

	return coreWallet.NextAccount(keyScope, name)
}

func (wallet *Wallet) KeyScope() *core.KeyScope {
	return nil
}

func getKeyScope() waddrmgr.KeyScope {
	return waddrmgr.KeyScopeBIP0044
}