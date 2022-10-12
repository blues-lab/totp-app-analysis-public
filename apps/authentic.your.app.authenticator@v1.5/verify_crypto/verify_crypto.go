package main

import (
	"encoding/base64"
	"fmt"

	"github.com/blues-lab/totp-app-analysis-public/utils"
)

const (
	password       = "1234567890"
	iterationCount = 65536
	keyBytesLength = 32
	ivBytesLength  = 16

	// The app used a static salt value
	salt = "ROYALEWITHCHEESEROYALEWITHCHEESE"

	// The backup Base64 value was saved to a file on the phone at path: /sdcard/Documents/AuthenticateBackups/<datetime>
	backupBase64 = "ZWhZ8TPtVRP/b6WX4A5ykljKjPourl0tC1/D2ytmAPWybt4TpDsHgoKhkxCc31D+aRFdv5OrS+FrBmEvOTvCDb6Qh0z/S4guKHIU5TNmEbuTL7IOrJhmx/YotsbgB0uG2LlC/Ve6Y7QlkMRkQhcupMv4t3T/4p7VUtDMuz5YYARpFHQVEWW0h7jR3WKskl5fZ4TYRjwHx4u4UODHMSjc9ef/rToZtkLg94ygnzYqExhXO2YdAPWogHfLO4mryZxw"
)

func main() {
	keyBytes := utils.Pkcs12KdfSha256(
		password,
		salt,
		iterationCount,
		keyBytesLength)

	backupBytes, err := base64.StdEncoding.DecodeString(backupBase64)
	if err != nil {
		panic(err)
	}

	// Note that the app does generate a random IV, but it is not used. This is
	// likely because the developer intended to use AES-CBC, but actually used
	// AES-ECB because of UX flaws in the Java crypto API. This is also likely
	// because this app is likely an unauthorized copy of
	// com.pixplicity.auth@v1.0.6, which has the same mistake.
	ivBytes := backupBytes[:ivBytesLength]
	fmt.Printf("ivBytes = %x\n\n", ivBytes)

	ciphertextBytes := backupBytes[ivBytesLength:]
	fmt.Printf("ciphertextBytes = %x\n\n", ciphertextBytes)

	plaintext := utils.DecryptAesEcbPaddedWithPKCS7(ciphertextBytes, keyBytes)
	fmt.Println("plaintext = ", string(plaintext))
}
