package main

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	keyBytesLength   = 32
	nonceBytesLength = 24
	numThreads       = 1

	password = "rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr"

	// Copy/pasted from network traffic with JSON field
	// `DUO_SECRET_BACKUP_CRYPTO_PARAMS`.
	saltOpsMemAlg = "8l5O00j8c6UOdrnpt6KQpA==:6:134217728:1"
)

func decryptOtpSecret(otpSecret string, keyBytes []byte) []byte {
	parts := strings.Split(otpSecret, ":")
	ciphertextBase64 := parts[0]
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("ciphertext as hex = %x\n", ciphertextBytes)

	nonceBase64 := parts[1]
	nonceBytes, err := base64.StdEncoding.DecodeString(nonceBase64)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("nonce as hex = %x\n", nonceBytes)

	// Convert data type from []byte to [32]byte required by secretbox.Open().
	var key [keyBytesLength]byte
	copy(key[:], keyBytes[:keyBytesLength])

	// Convert data type from []byte to [24]byte required by secretbox.Open().
	var nonce [nonceBytesLength]byte
	copy(nonce[:], nonceBytes[:nonceBytesLength])

	// Decrypt TOTP backup ciphertext.
	plaintext, ok := secretbox.Open(
		nil,
		ciphertextBytes,
		&nonce,
		&key,
	)
	if !ok {
		panic("failed to open ")
	}

	return plaintext
}

func deriveKeyFromPassword(password string, saltOpsMemAlg string) []byte {
	// fmt.Printf("password as hex = %x\n", passwordBytes)

	parts := strings.Split(saltOpsMemAlg, ":")
	saltBase64 := parts[0]
	saltBytes, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("salt as hex = %x\n", saltBytes)

	opsLimitAkaTime, _ := strconv.Atoi(parts[1])
	memLimitNumBytes, _ := strconv.Atoi(parts[2])
	memLimitNumKb := memLimitNumBytes / 1024
	// fmt.Println("memLimitNumKb = " + strconv.Itoa(memLimitNumKb))
	// memLimitNumMb := memLimitNumKb / 1024
	// fmt.Println("memLimitNumMb = " + strconv.Itoa(memLimitNumMb))
	// algorithm, _ := strconv.Atoi(parts[3])

	keyBytes := argon2.Key(
		[]byte(password),
		saltBytes,
		uint32(opsLimitAkaTime),
		uint32(memLimitNumKb),
		uint8(numThreads),
		uint32(keyBytesLength),
	)

	// fmt.Printf("key as hex = %x\n", keyBytes)
	return keyBytes
}

func main() {
	keyBytes := deriveKeyFromPassword(password, saltOpsMemAlg)

	// Copy/pasted from network traffic with JSON field `otpSecret`.
	ciphertextsNoncesBase64 := []string{
		"yxmVNJGRbOTwgAmv4FjgkUCU64EJLI3boQqwqCT52a1n9TqDPRv6ZQlqPLGO/6UyPVryTqZKL5rtsrULo9hbZkB0JEc\u003d:gcxgZEAylGM0klakmikgq4LNyE7yyaQu",
		"S8EkGtPPVOxW851ZSHLX3hQ0Jza7gdrCxviNf3A4u401wz6ZoNxHMPbue2Pkuazo8YhqNjI0NzJzEokLxEooxC1RBGk\u003d:w1U1xivtdSPJSZj+SDa8KZ79Et5E4Q6s",
	}

	for i, ciphertextNonceBase64 := range ciphertextsNoncesBase64 {
		plaintextBytes := decryptOtpSecret(ciphertextNonceBase64, keyBytes)
		fmt.Printf("Base32 TOTP secret #%d: %s\n", i+1, string(plaintextBytes))
	}
}
