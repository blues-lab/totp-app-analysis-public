package main

import (
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/nacl/secretbox"
	"golang.org/x/crypto/scrypt"
)

const (
	password       = "Password123!"
	keyBytesLength = 32

	nonceBytesLength = 24
	saltBytesLength  = 16

	// Parameteres for scrypt observed in decompiled APK.
	cpuMemoryFactorN = 32768
	blockSizeR       = 20
	parallelization  = 1

	// Copy/pasted from the network traffic, field name `backup`.
	backupBytesBase64 = "SOhRm917OEvA9L44YDLIKBBIXK8GoMp04Gm2YDNoSiXiP_nR-rsuUi608lt1ZiwUgnqQgKyQHi2CVs4V-wZKkpJCer6puvmGMy89jZIRXgHCH-6_WJPNVyZfWyRoXXXZPNeL_s9HspqATGZiijGaKyEths8FsjPN2WX64k2FaiNq2PbX6x2_rsv5YTvysQM0kGImpus3VQhA6FLlI5nvRTazoZ9j-LhOes69ysZgpyLp4J9AwY_okUsVbuOv1QdKklU"
)

func main() {
	// Base64 decode backupBytesBase64.
	backupBytes, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(backupBytesBase64)
	if err != nil {
		panic(err)
	}

	// Parse the salt, nonce, and ciphertext from backupBytes.
	saltIndex := len(backupBytes) - saltBytesLength
	saltBytes := backupBytes[saltIndex:]
	nonceBytes := backupBytes[:nonceBytesLength]
	ciphertextBytes := backupBytes[nonceBytesLength:saltIndex]

	// Derive the key from the password using scrypt.
	keyBytes, err := scrypt.Key(
		[]byte(password),
		saltBytes,
		cpuMemoryFactorN,
		blockSizeR,
		parallelization,
		keyBytesLength,
	)
	if err != nil {
		panic(err)
	}

	// Change data types from []byte to specific length arrays required by
	// secretbox.Open().
	var key [32]byte
	copy(key[:], keyBytes[:32])

	var nonce [24]byte
	copy(nonce[:], nonceBytes[:24])

	// Decrypt the TOTP backup.
	plaintextBytes, ok := secretbox.Open(
		nil,
		ciphertextBytes,
		&nonce,
		&key,
	)

	if !ok {
		panic("decryption error")
	}

	fmt.Println("TOTP backup plaintext = ", string(plaintextBytes))
}
