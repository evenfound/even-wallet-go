package btc

import (
	"errors"
	"github.com/btcsuite/btcwallet/waddrmgr"
	btcdWallet "github.com/btcsuite/btcwallet/wallet"
	_ "github.com/btcsuite/btcwallet/walletdb/bdb"
	"github.com/evenfound/even-wallet-go/core"
	"os"
	"path/filepath"
	"time"
)

type Wallet struct {
	core.HDWallet
}

var (
	_ core.WalletInterface = (*Wallet)(nil)
	_ core.BTCDInterface   = (*btcdWallet.Wallet)(nil)
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

	if err != nil {
		return true, err
	}

	defer f.Close()

	f.Write(core.Encrypt([]byte(wallet.Phrase), wallet.Password))

	return newWallet, nil
}

func (wallet Wallet) Authorize() (core.BTCDInterface, error) {
	var loader = wallet.GetLoader()

	return loader.OpenExistingWallet([]byte(wallet.Password), false)
}

func (wallet Wallet) Exists() (bool, error) {
	return wallet.GetLoader().WalletExists()
}

func (wallet Wallet) NextAccount(name string) (uint32, error) {

	var w, err = wallet.Authorize()

	if err != nil {
		return 0, err
	}

	// Converting to real type to get a real interface
	coreWallet, ok := w.(btcdWallet.Wallet)

	if !ok {
		return 0, errors.New(core.UnknownError)
	}

	var keyScope = getKeyScope()

	return coreWallet.NextAccount(keyScope, name)
}

func getKeyScope() waddrmgr.KeyScope {
	return waddrmgr.KeyScopeBIP0044
}

func (wallet Wallet) KeyScope() *core.KeyScope {
	return nil
}
