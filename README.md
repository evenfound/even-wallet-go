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
| foxnut/go-hdwallet | https://github.com/foxnut/go-hdwallet/blob/master/README.md |
| tyler-smith/go-bip39 | https://github.com/tyler-smith/go-bip39/blob/master/README.md |


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
   // generating a new seed / wallet
  var seed = wallet.Create("SEED_PHRASE", "STRONG_PASSWORD")
  
  // authorizing wallet to validate seed 
  
  if wallet.Authorize(seed) {
  		// generating a new address
  		var address  = wallet.NewAddress(0, 0, 0, 1)
  	} else {
  		fmt.Println(wallet.GetErrorMessage())
  	}
}


```


