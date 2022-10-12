package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"main/utils"
	"os"

	CryptoUtils "github.com/blues-lab/totp-app-analysis-public/utils"
)

const (
	password = "rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr"

	// The contents of the backup file created by the app's backup mechanism.
	backupFilename = "../../logs/2022-05-22T18-22-28Z-review-file-and-sharing/TOTP_Backup_1653243971.encrypt"
)

func main() {
	// The app uses a static IV of all zeroes.
	ivBytes := bytes.Repeat([]byte{0}, 16)

	// A single round of SHA256 is used as a KDF.
	h := sha256.New()
	h.Write([]byte(password))
	keyBytes := h.Sum(nil)

	ciphertextBase64, err := os.ReadFile(backupFilename)
	if err != nil {
		panic(err)
	}
	ciphertextBytes, _ := base64.StdEncoding.DecodeString(string(ciphertextBase64))

	// Decrypt the TOTP backup using AES-CBC.
	backupPlaintextBytes := CryptoUtils.DecryptAesCbcPaddedWithPKCS7(
		ciphertextBytes,
		keyBytes,
		ivBytes)
	fmt.Println("Plaintext Backup = ", string(backupPlaintextBytes))

	fmt.Println("Note that the totpSecrets are included in the backup as hex.")

	// Find, decode, and print the hex totp secrets within the plaintext backup.
	utils.FindDecodeAndPrintHexTotpSecrets(backupPlaintextBytes)
}
