# even-wallet-go
Hierarchical Deterministic Wallet based on golang


even-wallet-go  is a simple,flexible package which provides some methods to make and process hd wallets (BIP32,BIP39,BIP44).
We are using this package in our mobile apps based on Flutter to have a high performance and functionality. 

### Installation

```sh
$ go get github.com/evenfound/even-wallet-go
```

### Plugins

even-wallet-go is currently extended with the following plugins.

| Plugin | README |
| ------ | ------ |
| btcsuite/btcwallet | https://github.com/btcsuite/btcwallet/blob/master/README.md |
| ltcsuite/ltcswallet | https://github.com/ltcsuite/ltcwallet/blob/master/README.md |


### Development

For first be sure that you have already installed `JDK`,`JRE` (also `NDK`,` LLDB`, `CMake` packages)

```ssh
# installing and building gombile package
go get golang.org/x/mobile/cmd/gomobile

# initialize gombile
gomobile init

# go to project directory
cd ~/PATH_TO_YOUR_PROJECT

# make .aar package file for Flutter
gomobile bind -target=android .
```

### Usage
```
package main

import (
	"hdgen"
	"hdgen/core"
)

func main() {
  // initializing HDWallet to make available all functionality 
	var wallet = hdgen.Wallet{
		HDWallet: core.HDWallet{
			Name:     "wallet_name",
			Password: "strong_password",
			Phrase:   "PRIVATE_MNEMONIC_PHRASE",
			TestNet:  false,
		},
		Coin: "BTC",
	}

  // Creating new wallet
	wallet.Create()
}


```


