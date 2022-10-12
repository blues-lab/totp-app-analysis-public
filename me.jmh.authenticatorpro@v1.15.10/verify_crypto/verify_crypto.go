package main

import (
	"bytes"
	"crypto/aes"
	"crypto/sha1"
	"fmt"
	"os"

	"github.com/blues-lab/totp-app-analysis-public/utils"
	"golang.org/x/crypto/pbkdf2"
)

const (
	backupFilename = "backup-2022-05-20_223957.authpro"
	password       = "rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr"

	iterationCount  = 64000
	keyBytesLength  = 32
	saltBytesLength = 20
)

func main() {
	// Read backup from file.
	backupBytesAndHeader, err := os.ReadFile(backupFilename)
	if err != nil {
		panic(err)
	}

	// Remove header from the backup file.
	header := []byte("AuthenticatorPro")
	var backupBytes []byte
	if bytes.HasPrefix(backupBytesAndHeader, header) {
		fmt.Println("Header found and removed from backup file.")
		backupBytes = backupBytesAndHeader[len(header):]
	} else {
		fmt.Println("No header was found in the backup file. No action taken.")
		backupBytes = backupBytesAndHeader
	}

	// Split backupBytes into salt, iv, and ciphertext.
	saltBytes := backupBytes[:saltBytesLength]
	ivBytes := backupBytes[saltBytesLength : saltBytesLength+aes.BlockSize]
	ciphertextBytes := backupBytes[saltBytesLength+aes.BlockSize:]

	// Derive key from password using PBKDF2-HMAC-SHA1.
	keyBytes := pbkdf2.Key(
		[]byte(password),
		saltBytes,
		iterationCount,
		keyBytesLength,
		sha1.New)

	// Decrypt the TOTP backup.
	plaintextBytes := utils.DecryptAesCbcPaddedWithPKCS7(ciphertextBytes, keyBytes, ivBytes)
	fmt.Println("TOTP backup plaintext = ", string(plaintextBytes))
}
