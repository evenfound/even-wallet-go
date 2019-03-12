package core

import (
	"os"
	"path/filepath"
)

type BTCDInterface interface{}

type HDWallet struct {
	Name     string // wallet name
	Password string // password for symmetric encryption and wallet authorization
	Phrase   string // bip32 seed phrase
	TestNet  bool   // network type

	dataDirectory string // wallet directory
}

type WalletInterface interface {
	// Create function creates a new wallet (wallet.db) and
	// private.key file where stored encrypted privphrasee
	Create() (BTCDInterface, error)
	// Authorize function authorize's the wallet
	// using pubpass(password)
	Authorize() (BTCDInterface, error)
	// Exists function checks if the wallet already exists.
	// Check occurs using wallet name
	Exists() (bool, error)
}

// SetPath function sets the path where will be stored wallet
func (wallet *HDWallet) SetPath(path string) {
	wallet.dataDirectory = path
}

// Getting wallet path
func (wallet HDWallet) Path() string {

	var dbPath = wallet.dataDirectory

	var walletPath = os.Getenv(WalletPathKey)

	return filepath.Join(dbPath, walletPath, wallet.Name)
}
