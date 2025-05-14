package main

import (
	"fmt"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/base58"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/gagliardetto/solana-go"
	"github.com/tyler-smith/go-bip39"
)

func Generate() {
	// Generate 256 bits of entropy for a 24-word mnemonic
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		fmt.Printf("Error generating entropy: %v\n", err)
		return
	}

	// Generate mnemonic from entropy
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		fmt.Printf("Error generating mnemonic: %v\n", err)
		return
	}

	// Generate seed from mnemonic (no passphrase for simplicity)
	seed := bip39.NewSeed(mnemonic, "")

	// Derive a Solana private key using BIP-44 path m/44'/501'/0'/0'
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		fmt.Printf("Error creating master key: %v\n", err)
		return
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
			fmt.Printf("Error deriving key at index %d: %v\n", index, err)
			return
		}
	}

	// Get the private key (32 bytes) and ensure it's 64 bytes for Solana
	privateKeyBytes, err := key.ECPrivKey()
	if err != nil {
		fmt.Printf("Error getting private key: %v\n", err)
		return
	}
	// Solana expects a 64-byte private key (32-byte secret + 32-byte public key)
	// We derive the public key to construct the full 64-byte key
	pubKey, err := key.ECPubKey()
	if err != nil {
		fmt.Printf("Error getting public key: %v\n", err)
		return
	}
	pubKeyBytes := pubKey.SerializeCompressed()
	// Combine private key (32 bytes) and public key (32 bytes, padded if necessary)
	solanaPrivateKey := make([]byte, 64)
	copy(solanaPrivateKey[:32], privateKeyBytes.Serialize())
	copy(solanaPrivateKey[32:], pubKeyBytes[:32]) // Ensure 32 bytes

	// Create Solana keypair
	privateKey := solana.PrivateKey(solanaPrivateKey)
	publicKey := privateKey.PublicKey()

	// Print wallet details
	fmt.Println("=== Solana Wallet Details ===")
	fmt.Printf("Mnemonic: %s\n", mnemonic)
	fmt.Printf("Seed (hex): %x\n", seed)
	fmt.Printf("Private Key (base58): %s\n", base58.Encode(solanaPrivateKey))
	fmt.Printf("Public Key (base58): %s\n", publicKey.String())
}
func main() {
	// GetWalletAddressFromMnemonic("lobster order pen culture off fold fire brisk noble key organ chat private apple small cheap struggle intact phrase model blanket wagon throw enough", "")
	// GetWalletAddressFromMnemonic("rate skate globe bitter reduce reward festival crowd engage conduct dash prize urge boil drop tired noodle relax leopard chest estate disagree tunnel youth", "")
	GetWalletAddressFromMnemonic("mask dose stomach client upgrade fluid loan hard journey sniff paper river", "")

}
