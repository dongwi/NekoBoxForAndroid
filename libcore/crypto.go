package libcore

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func Sha1(data []byte) []byte {
	sum := sha1.Sum(data)
	return sum[:]
}

func Sha256Hex(data []byte) string {
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:])
}

func decrypt(ciphertext string) string {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("cbc decrypt err:", err)
		}
	}()

	block, err := aes.NewCipher([]byte("龍騰龘龘欣欣家国gooooing"))
	if err != nil {
		return ""
	}

	ciphercode, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return ""
	}

	iv := ciphercode[:aes.BlockSize]
	ciphercode = ciphercode[aes.BlockSize:]

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphercode, ciphercode)

	plaintext := string(ciphercode)
	return plaintext[:len(plaintext)-int(plaintext[len(plaintext)-1])]
}
