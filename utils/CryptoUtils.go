package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func DecryptAesGcm(ciphertextBytes []byte, keyBytes []byte, ivBytes []byte) []byte {
	aesCipher, err := aes.NewCipher(keyBytes)
	check(err)
	aesGcm, err := cipher.NewGCM(aesCipher)
	check(err)
	plaintextBytes, err := aesGcm.Open(nil, ivBytes, ciphertextBytes, nil)
	check(err)

	return plaintextBytes
}

// From: https://gist.github.com/stupidbodo/601b68bfef3449d1b8d9
func removePaddingPKCS7(src []byte) ([]byte, error) {
	length := len(src)
	lastByte := src[length-1]
	// fmt.Printf("Last byte as character = [%s]\n", string(lastByte))
	lastByteAsInt := int(lastByte)
	// fmt.Println("Padding length determined to be =", lastByteAsInt)

	if lastByteAsInt > length {
		return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
	}

	return src[:(length - lastByteAsInt)], nil
}

func DecryptAesEcbPaddedWithPKCS7(ciphertextBytes []byte, keyBytes []byte) []byte {
	// Must be at least one full block.
	if len(ciphertextBytes) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// Length must be a multiple of the block size because CBC mode always works in whole blocks.
	if len(ciphertextBytes)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	aesBlockCipher, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("Length of Ciphertext = %d\n\n", len(ciphertextBytes))
	paddedPlaintextBytes := make([]byte, len(ciphertextBytes))
	blockSize := aesBlockCipher.BlockSize()

	for blockStart, blockEnd := 0, blockSize; blockStart < len(ciphertextBytes); blockStart, blockEnd = blockStart+blockSize, blockEnd+blockSize {
		aesBlockCipher.Decrypt(paddedPlaintextBytes[blockStart:blockEnd], ciphertextBytes[blockStart:blockEnd])
	}

	plaintextBytes, err := removePaddingPKCS7(paddedPlaintextBytes)
	if err != nil {
		panic(err)
	}
	return plaintextBytes
}

func DecryptAesCbcPaddedWithPKCS7(ciphertextBytes []byte, keyBytes []byte, ivBytes []byte) []byte {
	// fmt.Printf("Length of Ciphertext = %d\n\n", len(ciphertextBytes))

	// Must be at least one full block.
	if len(ciphertextBytes) < aes.BlockSize {
		panic("ciphertext too short")
	}

	// Length must be a multiple of the block size because CBC mode always works in whole blocks.
	if len(ciphertextBytes)%aes.BlockSize != 0 {
		panic("ciphertext is not a multiple of the block size")
	}

	aesCipher, err := aes.NewCipher(keyBytes)
	check(err)
	aesCbcDecrypter := cipher.NewCBCDecrypter(aesCipher, ivBytes)

	paddedPlaintextBytes := make([]byte, len(ciphertextBytes))
	aesCbcDecrypter.CryptBlocks(paddedPlaintextBytes, ciphertextBytes)

	// fmt.Printf("len(paddedPlaintextBytes) = %d\n\n", len(paddedPlaintextBytes))
	// fmt.Printf("paddedPlaintextBytes = %x\n\n", paddedPlaintextBytes)
	// fmt.Printf("paddedPlaintextBytes as string = %s\n\n", paddedPlaintextBytes)

	plaintextBytes, err := removePaddingPKCS7(paddedPlaintextBytes)
	if err != nil {
		panic(err)
	}
	// fmt.Printf("plaintextBytes as string = %s\n\n", plaintextBytes)
	return plaintextBytes
}
