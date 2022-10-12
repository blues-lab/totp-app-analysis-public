package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"

	"github.com/blues-lab/totp-app-analysis-public/utils"
)

const (
	// These constants were observed in the Authy source code in the
	// getPBKDF2Key() function within the deprecated class:
	// com.authy.authy.util.CryptoHelper.java.
	iterationCount = 1000
	keyBytesLength = 32

	// The password used when creating a backup.
	password = "123456"

	// Copy/pasted from the network traffic, field name `salt`. See the included example traffic.
	saltBase64 = "o4y1APLlFB87FG7mkNxptPBq30wVvNlU"

	// Copy/pasted from the network traffic, field name `encrypted_seed`. See the included example traffic.
	ciphertextBase64 = "OzMycwH3YS57mSDXpt/hPSWRe2IwdpoeR+WHtEq0xBf5w5xPQm5nf6m7Y5dnlKPomsyUVzmy2mN2OCGH63Y7/Q=="
)

func main() {
	// Derive key from the password using PBKDF2-HMAC-SHA1.
	// Note that the app does not Base64 decode the salt.
	keyBytes := pbkdf2.Key(
		[]byte(password),
		[]byte(saltBase64),
		iterationCount,
		keyBytesLength,
		sha1.New)

	// The app uses a static IV of all zeroes.
	ivBytes := bytes.Repeat([]byte{0}, 16)
	// fmt.Printf("ivBytes as hex = %x\n", ivBytes)

	// Decode the Base64 ciphertext.
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		panic(err)
	}

	// Decrypt the ciphertext.
	plaintextBytes := utils.DecryptAesCbcPaddedWithPKCS7(ciphertextBytes, keyBytes, ivBytes)
	fmt.Printf("Plaintext TOTP secret in Base32 = %s\n", strings.ToUpper(string(plaintextBytes)))
}
