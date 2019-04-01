package hdgen

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// To see all test result in html run following command
// go test -coverprofile=coverage.out && go tool cover -html=coverage.out

const validMnemonicPhrase = "crunch ethics because blade cross remind accident buyer emotion pact double either chase myth drill"
const invalidMnemonicPhrase = "44xx44 emotion pact double either chase myth drill"

const validSeed = "13748d0c0b4ee48c124d50a8c6900d0f6041e3c6de756aeda1e7cdb2c7862ac6d7352d5e51fa91c2b9b6bd3ca77335c4be206da69c130ed653007b96a0c68f4d"
const invalidSeed1 = "1xx3748d0c0b4ee48c124d50a8c6900d0f6041e3c6de756aeda1e7cdb2c7862ac6d7352d5e51fa91c2b9b6bd3ca77335c4be206da69c130ed653007b96a0c68f4d"

const wif0001 = "5JVuyAvQ1cE6qQzzMb9PGQVCdyyPKLm7hn5dQ5Ud5H2wD6HNcAB"

func TestHDWallet_Create(t *testing.T) {
	Convey("Should return empty string and error message", t, func() {
		var hdWallet = HDWallet{}
		var testCases = []struct {
			mnemonic string
			password string

			seedHash     string
			errorMessage string
		}{
			{
				mnemonic: validMnemonicPhrase,
				password: "A123DSA321",

				seedHash:     validSeed,
				errorMessage: "",
			},
			{
				mnemonic: invalidMnemonicPhrase,
				password: "A123DSA321",

				seedHash:     "",
				errorMessage: "Invalid mnenomic",
			},
		}

		for _, test := range testCases {
			var message = fmt.Sprintf("Seed : %v", test.mnemonic)
			Convey(message, func() {
				var seed = hdWallet.Create(test.mnemonic, test.password)
				So(seed, ShouldEqual, test.seedHash)
				So(hdWallet.GetErrorMessage(), ShouldEqual, test.errorMessage)
			})
		}
	})
}

func TestHDWallet_Authorize(t *testing.T) {
	Convey("Testing wallet authorization", t, func() {
		var hdWallet = HDWallet{}
		var testCases = []struct {
			seedHash     string
			status       bool
			errorMessage string
		}{
			{
				seedHash:     invalidSeed1,
				status:       false,
				errorMessage: "encoding/hex: invalid byte: U+0078 'x'",
			},
			{
				seedHash:     validSeed,
				status:       true,
				errorMessage: "",
			},
		}

		for _, test := range testCases {
			var message = fmt.Sprintf("Giving seed hash to check authorization: Seed %v", test.seedHash)
			Convey(message, func() {
				var ok = hdWallet.Authorize(test.seedHash)
				So(ok, ShouldEqual, test.status)
				So(hdWallet.GetErrorMessage(), ShouldEqual, test.errorMessage)
			})
		}
	})
}

func TestHDWallet_IsTestNet(t *testing.T) {
	Convey("Given all necessary parameters to generate testnet address", t, func() {
		var wallet = HDWallet{}
		wallet.IsTestNet(true)
		wallet.Authorize(validSeed)

		var testCases = []struct {
			coinType     int
			accountIndex int
			changeIndex  int
			addressIndex int

			address      string
			errorMessage string
		}{
			{
				coinType:     0,
				accountIndex: 0,
				changeIndex:  0,
				addressIndex: 0,

				address:      "n1BvSGU5X6h9LuecDjP48nTYofmGjN99zP",
				errorMessage: "",
			}, {
				coinType:     0,
				accountIndex: 1,
				changeIndex:  0,
				addressIndex: 0,

				address:      "mmL4kYDvwRhw2NqDMAWSgjZ3AWY9C9XDfM",
				errorMessage: "",
			}, {
				coinType:     0,
				accountIndex: 0,
				changeIndex:  1,
				addressIndex: 0,

				address:      "mkaVTkGycoX4dZUqNyZtQ8cU672aJBNhRX",
				errorMessage: "",
			}, {
				coinType:     9999,
				accountIndex: 0,
				changeIndex:  0,
				addressIndex: 0,

				address:      "",
				errorMessage: UnsupportedCoinType,
			},
		}

		for _, test := range testCases {
			var message = fmt.Sprintf("coinType %v , Account %v , Change %v , Address %v",
				test.coinType, test.accountIndex, test.changeIndex, test.addressIndex)
			Convey(message, func() {
				var addr = wallet.NewAddress(test.coinType, test.accountIndex, test.changeIndex, test.addressIndex)
				So(addr, ShouldEqual, test.address)
				So(wallet.GetErrorMessage(), ShouldEqual, test.errorMessage)
			})
		}
	})
}

func TestHDWallet_NewAddress(t *testing.T) {
	var testCase = []struct {
		coinType     int
		accountIndex int
		changeIndex  int
		addressIndex int

		address      string
		errorMessage string
	}{
		{
			coinType:     0,
			accountIndex: 0,
			changeIndex:  0,
			addressIndex: 0,

			address:      "1Lfy9DP6i5FtZoAzWAQgJsFDwgAZtvtBCw",
			errorMessage: "",
		},
		{
			coinType:     0,
			accountIndex: 1,
			changeIndex:  0,
			addressIndex: 0,

			address:      "16p7TV8x8QGgFGMbdbY4rpLiJWwSEBqVa9",
			errorMessage: "",
		},
		{
			coinType:     0,
			accountIndex: 0,
			changeIndex:  1,
			addressIndex: 0,

			address:      "164YAhBzon5orT1DfQbWaDQ9E7RsQpp9Rs",
			errorMessage: "",
		},
		{
			coinType:     99999,
			accountIndex: 0,
			changeIndex:  0,
			addressIndex: 0,

			address:      "",
			errorMessage: UnsupportedCoinType,
		},
	}

	Convey("Given all necessary parameters to generate address", t, func() {
		var wallet = HDWallet{}
		wallet.Authorize(validSeed)
		for _, test := range testCase {
			var message = fmt.Sprintf("coinType %v , Account %v , Change %v , Address %v",
				test.coinType, test.accountIndex, test.changeIndex, test.addressIndex)
			Convey(message, func() {
				var addr = wallet.NewAddress(test.coinType, test.accountIndex, test.changeIndex, test.addressIndex)
				So(addr, ShouldEqual, test.address)
				So(wallet.GetErrorMessage(), ShouldEqual, test.errorMessage)
			})
		}
	})
}

func TestHDWallet_WIF(t *testing.T) {
	var testCases = []struct {
		coinType     int
		accountIndex int
		changeIndex  int
		addressIndex int

		wif          string
		errorMessage string

		shouldPanic bool

		authorized bool
	}{
		{
			coinType:     0,
			accountIndex: 0,
			changeIndex:  0,
			addressIndex: 0,

			wif:          "",
			errorMessage: UnAuthorized,
		},
		{
			coinType:     99999,
			accountIndex: 0,
			changeIndex:  0,
			addressIndex: 0,

			wif:          "",
			errorMessage: UnsupportedCoinType,

			authorized: true,
		},
		//{
		//	coinType:     0,
		//	accountIndex: 0,
		//	changeIndex:  0,
		//	addressIndex: 1,
		//
		//	wif:          wif0001,
		//	errorMessage: "",
		//
		//	authorized: true,
		//},
	}

	Convey("Testing WIF", t, func() {
		for _, test := range testCases {
			var wallet = HDWallet{}
			var message = fmt.Sprintf("coinType %v , Account %v , Change %v , Address %v , Authorized %v , WIF %v",
				test.coinType, test.accountIndex, test.changeIndex, test.addressIndex, test.authorized, test.wif)
			Convey(message, func() {
				if test.authorized {
					wallet.Authorize(validSeed)
				}
				var wif = func() string {
					return wallet.WIF(test.coinType, test.accountIndex, test.changeIndex, test.addressIndex)
				}
				if test.shouldPanic {
					So(wif, ShouldPanic)
				} else {
					So(wif(), ShouldEqual, test.wif)
					So(wallet.GetErrorMessage(), ShouldEqual, test.errorMessage)
				}
			})
		}
	})

}
