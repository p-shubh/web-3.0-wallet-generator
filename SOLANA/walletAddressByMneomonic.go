package main

import (
	"bytes"
	"crypto/ed25519"
	"fmt"

	"github.com/tyler-smith/go-bip39"
)

// GetWalletAddressFromMnemonic derives a Solana wallet address (public key) from a BIP-39 mnemonic.
// The passphrase is optional (use "" for no passphrase).
func GetWalletAddressFromMnemonic(mnemonic, passphrase string) (string, error) {
	// Generate seed from mnemonic and passphrase
	seed := bip39.NewSeed(mnemonic, passphrase)

	// Derive the master key from the seed
	// Replace with a proper implementation for deriving the master key
	// Generate a key pair using the seed
	publicKey, privateKey, err := ed25519.GenerateKey(bytes.NewReader(seed))
	if err != nil {
		return "", err
	}

	// Use the public key as the wallet address
	_ = privateKey // Private key can be stored securely if needed

	fmt.Println(fmt.Sprintf("%x", publicKey))

	// Encode the public key to a hexadecimal string
	return fmt.Sprintf("%x", publicKey), nil
}
