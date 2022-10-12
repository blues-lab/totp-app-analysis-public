package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/pbkdf2"

	"github.com/blues-lab/totp-app-analysis-public/utils"
)

const (
	password = "rRmhyojrhMyAFFD2_sEdXwiyJfTexRBr"

	// These constants were observed in the Authy source code in the
	// getPBKDF2Key() function within the deprecated class:
	// com.authy.authy.util.CryptoHelper.java.
	iterationCount   = 10000
	keyLengthInBytes = 32

	// Copy/pasted from the network traffic or file export, field is called
	// `servicesEncrypted`. You should copy the first (and only) value in the
	// array. Note: this value escapes forward slashes in the file export for some
	// reason. Therefore, it uses backticks so that golang interprets it literally
	// (i.e., does not try to interpret the escapes). See docs:
	// https://yourbasic.org/golang/multiline-string/#raw-string-literals

	// Value from backup to GDrive.
	// ciphertextSaltIvBase64WithEscapedForwardSlashes = `ukPMWA0Gnk4bwaxQ1W4pafOWA4m5LUNMDdZzsnGHIeyw2yWpN9yVEY+kgx4kv6NwZfB3ZBNqnvxZwtUwssiCgRMWLh7USA\/Ps7veQleoFioINMSsabejhxDb+zoU5tXeciI4oglq7ewdDK5oP5+dquDIHRmOBJ4fMSyvXx2hR+ZsipBXlMr5FrGWR7FCXvNX9pNIYny+t5BQTcgpaMhiF9zzatMqOS01o8az13UxfmWBJqf\/O7r0Vga39Cui1eOY3+5f0w3so46WEwuMHIolOSfZUOzVW068ZgNToFyzHsdV9TBhszEQDv5LzOlnA\/Z6Y9qeb+Bv7xXKyllFzG+GHfBPikFBXSFLzFwpBTy4+pUAbsQulG5yWj0YSEbHd2lLZsH0PoFq5LH4bBpHz9RuOGr\/3do5LB3UHW5qpym1gql4IeGMzYEZlk66o+iVmSoIUgrhu2APJ++0FrlXvp3b8cELm6+3tf1vgS+UsNqijmUIsQyAoqFia0zjeCjDbs22NX0JN8Y1SDQgReCbxFjqRH\/yINoVo4Wa+r+RY2eW8ZEhrvZYxgz2GcRjH\/f9nSilqSEsOMONDlEVAFv+zMV5tqwU3XXHdO0dwokYBJ9oRagq989CNfVNJ+mBm0GzmSIf+y6+jHhAphB7fBqWA8tWMZlz6CYnFtZrg8fAOp7jf9l5e9V8GyMQwoOhfHBIcArWaYlMtjMd94R\/oPlTbwoSHRmywTKa6HQC2yQyI97HJJmSEDdcWBzzBNHFUoUpcU0O\/hBJgGLyghoGAKQ+XNIU4aqckx39XUT4Dhi6xDVmhewjT74mjFkk4+XmwstK:vaV8r7OYOTtn54cbMzisxksgrCuAubhF1wI\/v\/k7+hNNcgDeU+ZKasPYzz4jBoyXwZ3KDNARPU+Q8ELXkIXfMDcOEGGeGPN+XQgmSjMKXo1njJcu8vYrqOukDWFVim4aZOBhCqWsQz1uvZfmQ7UAhCGN6EJtT41dOOlBTUqHL2DIalKz2oLTyLk+2gnrJW02X5OKXfu2TgA7Qu6jyQQazFhVowoqIa7G5oOnXMMaH8OXCxjtLyq3n8yA\/xOtm5MIktblET7tgLSkJ5zi\/Qd6AfofUeb3k2y25T5H91dztUQTvghIakAz6DhwUKN0t9u7JJWWuexYA0d9gOAXeked3Q==:cT5L699pMz3omNQQ`

	// Value from file export.
	ciphertextSaltIvBase64WithEscapedForwardSlashes = `6XKmOr7dKAJDF+u5tqthhZdxNJRBGpdh0KE07iK5A+RnhGELT8x5n1r0sz7ZfPoZT9bnN6N0lNI4WcUjhHhtMb\/jiduBlymiIg8QcDsC19G\/8hOrNM3tYHUjj68hWYrCl\/U+\/AtCCMIKUHyOsBenKJR9Io\/j8gakrYT8SCd2YZ9LF\/3sUCb1gLr3gTtt66YJ2qe5Frbmi6qtjnISNF+9ZyPyiXxoXqbVrXSR38U7Kxxq7+2rVQXRUKGElF+nVoe7j6jTrOkhlQ5h9DgbnK47i28eQgezhYvwFIGi4hsTPZnTkxK\/icz0qXVhQxR6DMymfAQeu80IirGVxAnp42hr+4rzKm9It9WLTjUFPgp1cL\/hFYIiMM1HZHJXtToMl\/\/7HooNGW8iZc\/sWhwS07zEM0I53PB1BcTvmtU8uJSJc66s8MP4qrV57xFfNl99kRGss8fjx+qbzxaiNEIpMPUWdqncyLvopFOTA8V9rl5vUXLWM7xQoETd17v3tyxhlfoQi4LlHxZ7M09YsNOiZTxlCranVcadnStwj+5\/3v8QONJL6Pq0zlZ63KqqZAFlOArsmosbrGaI6M7WEgnoeTotTKYAGKZtJkR2GFJKzTvTKKV0RJGLj\/6WZnXkyh5TynYlaqSBfXZKUIAeFvJUd5Sj+ps8nXnG2HPzzjKqAfJ96ovfovsZREpEo9pvGeXhh5u3Ep\/GMoE1JW4tVoUCvu5Dn9jLHZD\/kar0F22IgVHzMP0fjW9Foc7U0TD6tMe17BQLVMWCxiqm5fbCluJzJwsL1XuZcB6up3c10CIm1TXB5z9HxCF5foAwV2ZTIszItZ6+RKpqMi5vXRTxG0hLypdQGHS1SUX477C5k3xdX3IpZKbsB8mffVPXFxxNPJy3NEgygO2TLTnvjHu8EoA=:UeYAd6hoYjPbDttrameN5syPYxRTckiohUOOEbKdDPifTT1pH\/eXomV6DzGdsHG06k3MMyvsvl6SjpxnBDIF9rHeLcZYGrGxazEPqwgUZawF7kRPbMPxSDxub34AqfqTyMgnGbwtou07OaUSqULIvYVa4gXl2mIetEGqkz\/WWwefLxoUDt1eFLMjVl9X4EM5RUbbsIpJpl6NhyUA\/D3bTvp6xilKwvJjGHJuL9kqKBuTGu4D1jqFLAYWwD81lgXQf5+cN4RaHjk\/CCerxatpofgx98Ge6uzIz4LG+OqUUsEc294b6sZS+h2lmFk7h0SjUS0SkcgC+4ceqGGTtZAfWQ==:GXiSRZO52GO839xb`
)

func main() {
	// Unescape the forward slashes.
	escapedForwardSlash := "\\/"
	forwardSlash := "/"
	ciphertextSaltIvBase64 := strings.ReplaceAll(
		ciphertextSaltIvBase64WithEscapedForwardSlashes,
		escapedForwardSlash,
		forwardSlash)

	// Split apart the ciphertext, salt, and iv.
	parts := strings.Split(ciphertextSaltIvBase64, ":")
	ciphertextBase64 := parts[0]
	saltBase64 := parts[1]
	ivBase64 := parts[2]

	// Base64 decode the ciphertext, salt, and iv.
	ciphertextBytes, err := base64.StdEncoding.DecodeString(ciphertextBase64)
	if err != nil {
		panic(err)
	}

	saltBytes, err := base64.StdEncoding.DecodeString(saltBase64)
	if err != nil {
		panic(err)
	}

	ivBytes, err := base64.StdEncoding.DecodeString(ivBase64)
	if err != nil {
		panic(err)
	}

	// Derive the key from the password using PBKDF2-HMAC-SHA256.
	keyBytes := pbkdf2.Key(
		[]byte(password),
		saltBytes,
		iterationCount,
		keyLengthInBytes,
		sha256.New)

	// Decrypt the TOTP backup using AES-GCM.
	plaintextBytes := utils.DecryptAesGcm(ciphertextBytes, keyBytes, ivBytes)
	fmt.Printf("TOTP backup plaintext = %s\n\n", plaintextBytes)
}
