package main

import (
	"encoding/base64"
	"fmt"

	"github.com/blues-lab/totp-app-analysis-public/utils"
)

const (
	password       = "rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr"
	iterationCount = 65536
	keyBytesLength = 32
	ivBytesLength  = 16

	// The app used a static salt value
	salt = "ROYALEWITHCHEESEROYALEWITHCHEESE"

	// Copy/pasted from the backup created by the app. The easiest way I found to
	// get this value off the Android phone was to email it to myself (the app
	// uses the standard Android sharing menu).
	backupBase64 = "slSb6K7Gc73RgnRbOt0kcLV4NtirU7PR8kmEjeLw5dMqyQoYHD4Nag/GUc6p8d8zKetPwOx2CBjyy/B0FJMsdahxtZVw7XiLY6SS+mNWv7WpL+eAYCel+fxdbRqL2qmBlWXIfcIEIUIn6IWzUzFo5vmzWW2b5M3qWpYVJdgm133R1EWoJd3Qq0+YeuMMWA6NZHmYFH7sRcTjCm/YlLyN/j4Ja20jGhu1ciY4kcT6+NedgfbGYs2RNxl5ZbqXFjHc"
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
	// AES-ECB because of UX flaws in the Java crypto API. This issue was
	// disclosed to the developer and fixed in version 1.2.4 of the app. You can
	// see the fixed app using AES-CBC in directory com.pixplicity.auth@v1.2.4.
	ivBytes := backupBytes[:ivBytesLength]
	fmt.Printf("ivBytes = %x\n\n", ivBytes)

	ciphertextBytes := backupBytes[ivBytesLength:]
	fmt.Printf("ciphertextBytes = %x\n\n", ciphertextBytes)

	plaintext := utils.DecryptAesEcbPaddedWithPKCS7(ciphertextBytes, keyBytes)
	fmt.Println("plaintext = ", string(plaintext))
}
