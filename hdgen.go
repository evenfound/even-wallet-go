package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/ripemd160"
)

type keychain struct {
	masterPrivateKey string
	masterPublicKey  string
}

type HDWallet struct {
	seed []byte
	keychain
}

func main() {

	var hdwallet = new(HDWallet)

	hdwallet.GenerateWallet(WalletGenerator{
		"wallet1",
		[]string{"yuri", "gasparyan", "ashot", "erevan", "massiv", "nor", "norq", "tun", "avto", "artcoding", "dantser", "even"},
		"880088aa",
	})
	hdwallet.GenerateAddress(10, 10)
}

// Create a new HD wallet
func (hd *HDWallet) GenerateWallet(wg WalletGenerator) (*HDWallet, error) {

	err := wg.validate()

	if err != nil {
		return nil, err
	}

	hd.seed = bip39.NewSeed(wg.toString(), "")

	masterPrivateKey, err := bip32.NewMasterKey(hd.seed)

	if err != nil {
		return nil, err
	}

	hd.masterPrivateKey = masterPrivateKey.String()

	var masterPublicKey = masterPrivateKey.PublicKey()

	hd.masterPublicKey = masterPublicKey.String()

	wg.createDataDir()

	return hd, nil
}

// Generate Address
func (hd *HDWallet) GenerateAddress(x, y int) {
	var xpubByte = []byte(hd.masterPublicKey)

	var hasher = sha256.New()

	hasher.Write(xpubByte)

	sha := hex.EncodeToString(hasher.Sum(nil))

	var ripemdHasher = ripemd160.New()

	ripemdHasher.Write([]byte(sha))

	ripemdHash := hex.EncodeToString(ripemdHasher.Sum(nil))

	fmt.Println(hd.masterPublicKey, sha, ripemdHash)

}

// Get list of accounts
func (hd *HDWallet) GetAccounts() {}

// Delete account by hash
func (hd *HDWallet) DeleteAccount() {

}
