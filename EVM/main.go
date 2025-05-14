// package main

// func main() {

// 	GenerateAccount()

// }

package main

import (
	"fmt"
	"log"
	"github.com/miguelmota/go-ethereum-hdwallet"
)

func generateWallet(mnemonic string, path string) {
	// Create wallet from mnemonic
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the derivation path
	parsedPath := hdwallet.MustParseDerivationPath(path)

	// Derive the account from the wallet
	account, err := wallet.Derive(parsedPath, false)
	if err != nil {
		log.Fatal(err)
	}

	// Get private key in hex
	privateKeyHex, err := wallet.PrivateKeyHex(account)
	if err != nil {
		log.Fatal(err)
	}

	// Print wallet address and private key
	fmt.Println("Wallet Address:", account.Address.Hex())
	fmt.Println("Private Key:", privateKeyHex)
}

func main() {
	// Generate new mnemonic
	mnemonic, err := hdwallet.NewMnemonic(128)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Generated Mnemonic:", mnemonic)

	// List of common cryptocurrency derivation paths
	paths := []struct {
		crypto   string
		derivationPath string
	}{
		{"Ethereum (ETH)", "m/44'/60'/0'/0/0"},
		{"Binance Smart Chain (BSC)", "m/44'/60'/0'/0/0"},
		{"Polygon (MATIC)", "m/44'/60'/0'/0/0"},
		{"Bitcoin (BTC)", "m/44'/0'/0'/0/0"},
		{"Litecoin (LTC)", "m/44'/2'/0'/0/0"},
		{"Dogecoin (DOGE)", "m/44'/3'/0'/0/0"},
		{"Bitcoin Cash (BCH)", "m/44'/145'/0'/0/0"},
		{"Zcash (ZEC)", "m/44'/133'/0'/0/0"},
		{"Ripple (XRP)", "m/44'/144'/0'/0/0"},
		{"Monero (XMR)", "m/44'/128'/0'/0/0"}, // Custom path for Monero
	}

	// Generate wallets for each cryptocurrency path
	for _, path := range paths {
		fmt.Printf("\nGenerating wallet for %s...\n", path.crypto)
		generateWallet(mnemonic, path.derivationPath)
	}
}
