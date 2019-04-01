package hdgen

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

var transaction = Transaction{}

const signedTransaction = "0100000001f341747f616f6c208057c2090cc500f407c0f4ca7e64e9bb051132d2a89fb922000000008a47304402200b16503a823d1e580a4ab3dbd872fb20c560c3278df25418b363febe0dd8698c022076d446d4150d2eaa5da68158403570a543df91c01c1bd309a970b6b1e93a1e24014104e03bd1fa5a56ef0bc9d7b32c7b06148003e1389982aba6521a01bbfcb51392df45116a7def38ef7201640d24e4a5fbb56e9c6905cc1c692b9486d08a9e17c164ffffffff0101000000000000001976a9143fc191a0758036a07b6cd5b184d0f572ba83f14488ac00000000"

const unsignedTransaction = "0100000001b3823c6611dd01f2a6b33ed6ff8e9dea53af06f740bf66dac7be15d254ca5b4f0000000000ffffffff0101000000000000001976a9147f052efde790908c639f6f3c2d73183592e106d388ac00000000"

const validTransactionHash = "4f5bca54d215bec7da66bf40f706af53ea9d8effd63eb3a6f201dd11663c82b3"

const invalidTransactionHash = "invalid_transaction_hash"

const invalidWIF = "invalid_wif"

const invalidDestinationAddress = "invalid_address"

const unsupportedCoinType = 999

const validAddress = "16p7TV8x8QGgFGMbdbY4rpLiJWwSEBqVa9"


func TestTransaction_Make(t *testing.T) {
	Convey("Should receive an error", t, func() {
		var transaction = Transaction{}
		transaction.Make(unsupportedCoinType, wif0001, validAddress, validTransactionHash, 1)
		Convey("When coin type is supported", func() {
			So(transaction.GetErrorMessage(), ShouldEqual, UnsupportedCoinType)
		})
	})
}

func TestTransaction_Make2(t *testing.T) {
	Convey("Should receive a panic", t, func() {
		Convey("When transaction hash is invalid", func() {
			var transaction = Transaction{}
			make := func() {
				transaction.Make(0, wif0001, validAddress, invalidTransactionHash, 1)
			}
			So(make, ShouldPanic)
		})
	})
}

func TestTransaction_Make3(t *testing.T) {
	Convey("Should receive a panic", t, func() {
		Convey("When destination address is not valid", func() {
			var transaction = Transaction{}
			make := func() {
				transaction.Make(0, wif0001, invalidDestinationAddress, validTransactionHash, 1)
			}
			So(make, ShouldPanic)
		})
	})
}

func TestTransaction_Make4(t *testing.T) {
	Convey("Should receive an error", t, func() {
		Convey("When WIF is not valid", func() {
			var transaction = Transaction{}
			transaction.Make(0, invalidWIF, validAddress, validTransactionHash, 1)
			So(transaction.GetErrorMessage(), ShouldEqual, "malformed private key")
		})
	})
}

func TestTransaction_Make5(t *testing.T) {
	Convey("Should return signed and unsigned transactions", t, func() {
		var transaction = Transaction{}
		Convey("When passed arguments is ok", func() {
			transaction.Make(0, wif0001, validAddress, validTransactionHash, 1)
			So(transaction.GetSignedTransaction(), ShouldEqual, signedTransaction)
			So(transaction.GetUnsignedTransaction(), ShouldEqual, unsignedTransaction)
		})
	})
}

func TestTransaction_GetErrorMessage(t *testing.T) {
	Convey("Should receive an error", t, func() {
		transaction.Make(unsupportedCoinType, wif0001, validAddress, validTransactionHash, 1)
		Convey("When coin type is supported", func() {
			So(transaction.GetErrorMessage(), ShouldEqual, UnsupportedCoinType)
		})
	})
}
