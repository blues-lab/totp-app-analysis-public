package utils

import (
	"encoding/hex"
	"fmt"
	"regexp"
)

func FindDecodeAndPrintHexTotpSecrets(backupPlaintextBytes []byte) {
	r, _ := regexp.Compile("[0-9A-F]{64}")
	totpSecretsHex := r.FindAllString(string(backupPlaintextBytes), -1)

	for i, totpSecretHex := range totpSecretsHex {
		totpSecretAscii, err := hex.DecodeString(totpSecretHex)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Totp secret #%d: %s\n", i+1, totpSecretAscii)
	}
}
