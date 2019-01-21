package main

import (
	"errors"
	"os"
	"os/user"
	"regexp"
	"strings"
)

type WalletGenerator struct {
	Name     string
	Phrase   []string
	Password string
}

var (
	dataDir string
)

// Convert phrase slice to string
func (wg WalletGenerator) toString() string {
	return strings.Join(wg.Phrase, " ")
}

// Validate phrase
func (wg WalletGenerator) validate() (err error) {

	if len(wg.Phrase) < 12 || len(wg.Phrase) > 24 {
		err = errors.New("Phrase must contains [12,24] words")
	}

	phraseStr := wg.toString()

	isOk, _ := regexp.Match("^[a-zA-Z\\s]*$", []byte(phraseStr))

	if !isOk {
		err = errors.New("Phrase must contains only letters")
	}

	return
}


// Creating a directory
func (wg WalletGenerator) createDataDir() error {
	osUser, err := user.Current()

	if err != nil {
		return err
	}

	userDataDir := osUser.HomeDir + "\\even-daemon\\wallet"

	return os.Mkdir(userDataDir, 0777)
}
