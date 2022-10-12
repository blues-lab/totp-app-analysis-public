package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/blues-lab/totp-app-analysis-public/utils"
	"golang.org/x/crypto/pbkdf2"
)

const (
	keyBytesLength = 32
	password       = "0000"

	// Copy/pasted from the network traffic, field is called `hash_iterations`.
	iterationCount = 10000

	// Copy/pasted from the network traffic, field is called `hash_salt`.
	saltBase64 = "q/jhM8M5JxP9p4RsVbPsQg=="

	// Copy/pasted from the network traffic, field is called `initialization_vector`.
	ivBase64 = "bE30TO+NWJIx+NW7xHZz+g=="

	// Copy/pasted from the network traffic, field is called `encrypted_bundle`.
	totpBackupCiphertextBase64 = "I2tgOcAGCD4ftJ8qtOGFiysE0eS8Ule8190RH1xwhIJlrpr+NKfq8fLSj/7gJ/0WiIXXA5EyPjqsUUe2BOakMWsARFGYmCGDEK2R3ewx5WI8z8VTsKwJu1vmHirBUbPP36R+OupiFAW4kYpg08wFAlwygDFY3hkJWpiYk4zzRf81kO7TQNUgdBoXQtHWShaKExoI74FwymM/pulKIL3+soBaMLqNn7wxKmgQf/WBn4llR/Qg7PjdMWQyFcJoNfQ7z3UDVhpHggrrU02WoC+jMN/CsKSvUJp4/QpaMoqK76LKtDrhsUgwbYgQ0LnBDcqUL9qmlfNoXrYntWbdgIwVDdVs0j9tm5j30twxVfmZJMA="

	// Copy/pasted from the network traffic, field is called `encrypted_keys`.
	// You should copy the first (and only) value in the array.
	// Note: this value escapes forward slashes for some reason, so it uses backticks
	// so that golang interprets it literally (i.e., does not try to interpret the escapes).
	// See docs: https://yourbasic.org/golang/multiline-string/#raw-string-literals
	wrappedKeyCiphertextBase64WithEscapedForwardSlashes = `9dL3u2P5kzZozD9Jjp92EHxrPKr2jwrWWQm\/hfZzAvsuF\/9a37p9ue89IRMh\/gSR`
)

func main() {
	// Base64 decode the salt, iv, and totpBackup values.
	saltBytes, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		panic(err)
	}

	ivBytes, err := base64.StdEncoding.DecodeString(ivBase64)
	if err != nil {
		panic(err)
	}

	totpBackupCiphertext, err := base64.StdEncoding.DecodeString(totpBackupCiphertextBase64)
	if err != nil {
		panic(err)
	}

	// Unescape the forward slashes and Base64 decode the wrapped key ciphertext.
	escapedForwardSlash := "\\/"
	forwardSlash := "/"
	wrappedKeyCiphertextBase64 := strings.ReplaceAll(
		wrappedKeyCiphertextBase64WithEscapedForwardSlashes,
		escapedForwardSlash,
		forwardSlash)
	wrappedKeyCiphertext, err := base64.StdEncoding.DecodeString(wrappedKeyCiphertextBase64)
	if err != nil {
		panic(err)
	}

	// Derive the key from the password using PBKDF2-HMAC-SHA256.
	derivedKeyBytes := pbkdf2.Key(
		[]byte(password),
		saltBytes,
		iterationCount,
		keyBytesLength,
		sha256.New)

	// Decrypt the wrapped key.
	wrappedKeyPlaintextBase64 := utils.DecryptAesCbcPaddedWithPKCS7(wrappedKeyCiphertext, derivedKeyBytes, ivBytes)

	// Base64 decode the decrypted wrapped key.
	wrappedKeyPlaintext, _ := base64.StdEncoding.DecodeString(string(wrappedKeyPlaintextBase64))

	// Decrypt the totpBackup using the same key that encrypted the wrapped key.
	// This wrapping provides no additional security.
	totpBackupPlaintext := utils.DecryptAesCbcPaddedWithPKCS7(totpBackupCiphertext, wrappedKeyPlaintext, ivBytes)

	fmt.Println("TOTP backup plaintext = ", string(totpBackupPlaintext))
}
