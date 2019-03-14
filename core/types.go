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
	Coin     string // coin name

	dataDirectory string // wallet directory
	error         string // error message
}

type KeyScope struct {
	// Purpose is the purpose of this key scope. This is the first child of
	// the master HD key.
	Purpose uint32

	// Coin is a value that represents the particular coin which is the
	// child of the purpose key. With this key, any accounts, or other
	// children can be derived at all.
	Coin uint32
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

	NextAccount(name string) (uint32, error)

	//KeyScope() *KeyScope
}

// SetPath function sets the path where will be stored wallet
func (wallet *HDWallet) SetPath(path string) {
	wallet.dataDirectory = path
}

// Creating log file if not exists
func (wallet *HDWallet) Logger(fPath string) error {

	return nil
}

// Setting error message
func (wallet *HDWallet) SetErrorMessage(message string) {
	wallet.error = message
}

// Getting error message
func (wallet *HDWallet) GetErrorMessage() string {
	return wallet.error
}

// Getting wallet path
func (wallet HDWallet) Path() string {

	var dbPath = wallet.dataDirectory

	var walletPath = os.Getenv(WalletPathKey)

	return filepath.Join(dbPath, wallet.Coin, walletPath, wallet.Name)
}
