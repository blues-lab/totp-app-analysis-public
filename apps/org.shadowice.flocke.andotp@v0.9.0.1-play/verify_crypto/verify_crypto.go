package main

import (
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"os"

	"golang.org/x/crypto/pbkdf2"

	"github.com/blues-lab/totp-app-analysis-public/utils"
)

const (
	password       = "rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr"
	backupFilename = "otp_accounts_2022-05-20_19-38-55.json.aes"

	// Constants taken from org.shadowice.flocke.andotp.Utilities.Constants.java.
	// Note that saltLength is incorrectly listed as 16 bytes in Constants.java
	// (PBKDF2_SALT_LENGTH), but is clearly only 12 bytes when created
	// (org.shadowice.flocke.andotp.Utilities.BackupHelper.java#L105) and when read
	// from file (org.shadowice.flocke.andotp.Tasks.EncryptedRestoreTask.java#L42).
	iterationBytesLength = 4
	saltBytesLength      = 12
	ivBytesLength        = 12
	keyBytesLength       = 32
)

func main() {
	// Read bytes from backup file that was created by the app.
	backupBytes, err := os.ReadFile(backupFilename)
	if err != nil {
		panic(err)
	}

	// Parse the iteration count, salt, iv, and ciphertext from the backup.
	iterationCount := int(binary.BigEndian.Uint32(backupBytes[:iterationBytesLength]))
	saltBytes := backupBytes[iterationBytesLength : iterationBytesLength+saltBytesLength]
	ivBytes := backupBytes[iterationBytesLength+saltBytesLength : iterationBytesLength+saltBytesLength+ivBytesLength]
	ciphertextBytes := backupBytes[iterationBytesLength+saltBytesLength+ivBytesLength:]

	// Derive the key from the password using PBKDF2-HMAC-SHA1.
	keyBytes := pbkdf2.Key(
		[]byte(password),
		saltBytes,
		iterationCount,
		keyBytesLength,
		sha1.New)

	// Decrypt the TOTP backup using AES-GCM.
	plaintextBytes := utils.DecryptAesGcm(ciphertextBytes, keyBytes, ivBytes)
	fmt.Printf("TOTP backup plaintext = %s\n\n", plaintextBytes)
}
