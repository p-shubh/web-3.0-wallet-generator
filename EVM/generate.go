package main

import (
	"fmt"
	"log"

	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

func GenerateAccount() {
	// Generate a new 12-word mnemonic
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Generated Mnemonic:", mnemonic)

	// Create wallet
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	// Use 'm' as the root path (not "")
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}

	// Get private key
	privateKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Wallet Address:", account.Address.Hex())
	fmt.Println("Private Key:", privateKeyHex)
}
