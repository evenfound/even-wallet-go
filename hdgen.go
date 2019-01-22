package hdgen

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip39"
)

type Wallet struct {
	mnemonic    string
	path        string
	root        *hdkeychain.ExtendedKey
	extendedKey *hdkeychain.ExtendedKey
	privateKey  *ecdsa.PrivateKey
	publicKey   *ecdsa.PublicKey
}

type Config struct {
	Mnemonic  string
	Path      string
	TokenCode string
}

func GenerateWallet(config *Config) (*Wallet, error) {

	if config.Mnemonic == "" {
		return nil, errors.New("The mnemonic phrase required")
	}

	var seed = bip39.NewSeed(config.Mnemonic, "")

	dpath, err := accounts.ParseDerivationPath(config.Path)

	if err != nil {
		return nil, err
	}

	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)

	if err != nil {
		return nil, err
	}

	key := masterKey

	for _, n := range dpath {
		key, err = key.Child(n)
		if err != nil {
			return nil, err
		}
	}

	privateKey, err := key.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return nil, err
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("Failed to get public key")
	}

	wallet := &Wallet{
		mnemonic:    config.Mnemonic,
		path:        config.Path,
		root:        masterKey,
		extendedKey: key,
		privateKey:  privateKeyECDSA,
		publicKey:   publicKeyECDSA,
	}

	return wallet, nil
}

func (wallet Wallet) Derive(index interface{}) (*Wallet, error) {
	var idx uint32
	switch v := index.(type) {
	case int:
		idx = uint32(v)
	case int8:
		idx = uint32(v)
	case int16:
		idx = uint32(v)
	case int32:
		idx = uint32(v)
	case int64:
		idx = uint32(v)
	case uint:
		idx = uint32(v)
	case uint8:
		idx = uint32(v)
	case uint16:
		idx = uint32(v)
	case uint32:
		idx = v
	case uint64:
		idx = uint32(v)
	default:
		return nil, errors.New("unsupported index type")
	}

	fmt.Println(idx)

	address, err := wallet.extendedKey.Child(idx)
	if err != nil {
		return nil, err
	}

	privateKey, err := address.ECPrivKey()
	privateKeyECDSA := privateKey.ToECDSA()
	if err != nil {
		return nil, err
	}

	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("failed ot get public key")
	}

	path := fmt.Sprintf("%s/%v", wallet.path, idx)

	return &Wallet{
		path:        path,
		root:        wallet.extendedKey,
		extendedKey: address,
		privateKey:  privateKeyECDSA,
		publicKey:   publicKeyECDSA,
	}, nil
}

// PrivateKey ...
func (s Wallet) PrivateKey() *ecdsa.PrivateKey {
	return s.privateKey
}

// PrivateKeyBytes ...
func (s Wallet) PrivateKeyBytes() []byte {
	return crypto.FromECDSA(s.PrivateKey())
}

// PrivateKeyHex ...
func (s Wallet) PrivateKeyHex() string {
	return hexutil.Encode(s.PrivateKeyBytes())[2:]
}

// PublicKey ...
func (s Wallet) PublicKey() *ecdsa.PublicKey {
	return s.publicKey
}

// PublicKeyBytes ...
func (s Wallet) PublicKeyBytes() []byte {
	return crypto.FromECDSAPub(s.PublicKey())
}

// PublicKeyHex ...
func (s Wallet) PublicKeyHex() string {
	return hexutil.Encode(s.PublicKeyBytes())[4:]
}

// Address ...
func (s Wallet) Address() common.Address {
	return crypto.PubkeyToAddress(*s.publicKey)
}

// AddressHex ...
func (s Wallet) AddressHex() string {
	return s.Address().Hex()
}

// Path ...
func (s Wallet) Path() string {
	return s.path
}
