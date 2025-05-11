package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/cosmos/go-bip39"
	"github.com/gagliardetto/solana-go"
)

// GetWalletAddressFromMnemonic derives a Solana wallet address (public key) from a BIP-39 mnemonic.
// The passphrase is optional (use "" for no passphrase).
func GetWalletAddressFromMnemonic(mnemonic, passphrase string) (string, error) {
	// Validate the mnemonic
	if !bip39.IsMnemonicValid(mnemonic) {
		return "", fmt.Errorf("invalid mnemonic")
	}

	// Generate seed from mnemonic
	seed := bip39.NewSeed(mnemonic, passphrase)

	// Derive a Solana private key using BIP-44 path m/44'/501'/0'/0'
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return "", fmt.Errorf("error creating master key: %v", err)
	}

	// Derive the BIP-44 path: m/44'/501'/0'/0'
	path := []uint32{
		44 + hdkeychain.HardenedKeyStart,  // Purpose (44')
		501 + hdkeychain.HardenedKeyStart, // Coin type (Solana)
		0 + hdkeychain.HardenedKeyStart,   // Account (0')
		0,                                 // Change (0)
	}
	key := masterKey
	for _, index := range path {
		key, err = key.Child(index)
		if err != nil {
			return "", fmt.Errorf("error deriving key at index %d: %v", index, err)
		}
	}

	// Get the private key (32 bytes)
	privateKeyBytes, err := key.ECPrivKey()
	if err != nil {
		return "", fmt.Errorf("error getting private key: %v", err)
	}

	// Derive the public key (32 bytes, padded if necessary)
	pubKey, err := key.ECPubKey()
	if err != nil {
		return "", fmt.Errorf("error getting public key: %v", err)
	}
	pubKeyBytes := pubKey.SerializeCompressed()

	// Create a 64-byte Solana private key (32-byte secret + 32-byte public key)
	solanaPrivateKey := make([]byte, 64)
	copy(solanaPrivateKey[:32], privateKeyBytes.Serialize())
	copy(solanaPrivateKey[32:], pubKeyBytes[:32]) // Ensure 32 bytes

	// Create Solana keypair
	privateKey := solana.PrivateKey(solanaPrivateKey)
	publicKey := privateKey.PublicKey()

	fmt.Println("publicKey.String() : ", publicKey.String())

	// Return the public key (wallet address) in base58
	return publicKey.String(), nil
}
