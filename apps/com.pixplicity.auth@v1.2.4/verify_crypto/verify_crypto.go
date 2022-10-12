package main

import (
	"encoding/base64"
	"fmt"

	"github.com/blues-lab/totp-app-analysis-public/utils"
)

const (
	password        = "rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr"
	iterationCount  = 65536
	keyBytesLength  = 32
	saltBytesLength = 32
	ivBytesLength   = 16

	// Copy/pasted from the backup created by the app. The easiest way I found to
	// get this value off the Android phone was to email it to myself (the app
	// uses the standard Android sharing menu).
	backupBase64 = "QjaDIGiE65GcENRMtE06xCZHAPZa7Q2ajvwsgDjFuQi+S7iiDfRQ9ltOb8+MVI9nqsrYrJeTs2Wo3iWeFCaVTEzuDB3VQy4TzncDyQjtzVQ0OomqHinuu6HgLyUNed/uZE9dFEjGTT9ZBTTau15Ig3UfM2tIA5VDETJaD1araFxXOzuhhcsYjKzGYfzhu9MrpWpyCkq3I6T1ABoSlU/1k7rBcNCSwt6Ov2E7Qs6s45oqAZ60YMM5nXoJx/lG3n21wEj3Sd/5jcUBpHRoW4AzILrOckXRDlyeciyQlRafS8A="
)

func main() {
	backupBytes, err := base64.StdEncoding.DecodeString(backupBase64)
	if err != nil {
		panic(err)
	}

	saltBytes := backupBytes[:saltBytesLength]
	fmt.Printf("saltBytes = %x\n\n", saltBytes)

	ivBytes := backupBytes[saltBytesLength : saltBytesLength+ivBytesLength]
	fmt.Printf("ivBytes = %x\n\n", ivBytes)

	ciphertextBytes := backupBytes[saltBytesLength+ivBytesLength:]
	fmt.Printf("ciphertextBytes = %x\n\n", ciphertextBytes)

	keyBytes := utils.Pkcs12KdfSha256(
		password,
		string(saltBytes),
		iterationCount,
		keyBytesLength)

	plaintext := utils.DecryptAesCbcPaddedWithPKCS7(ciphertextBytes, keyBytes, ivBytes)
	fmt.Println("plaintext = ", string(plaintext))
}
