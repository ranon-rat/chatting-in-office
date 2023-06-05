package core

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
)

func InitCipher(rendezvous string) {
	var err error
	if Caes, err = aes.NewCipher([]byte(AddInfo(rendezvous))); err != nil {
		panic(err)
	}

}
func AddInfo(key string) string {
	if len(key) >= 16 {
		return key[:16]
	}
	for len(key) < 16 {
		key += " "
	}
	return key
}
func EncryptMessage(message string) (string, error) {
	byteMsg := []byte(message)
	block := Caes

	cipherText := make([]byte, aes.BlockSize+len(byteMsg))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", fmt.Errorf("could not encrypt: %v", err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], byteMsg)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func DecryptMessage(message string) (string, error) {
	cipherText, err := base64.StdEncoding.DecodeString(message)
	if err != nil {
		return "", fmt.Errorf("could not base64 decode: %v", err)
	}

	block := Caes

	if len(cipherText) < aes.BlockSize {
		return "", fmt.Errorf("invalid ciphertext block size")
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), nil
}
