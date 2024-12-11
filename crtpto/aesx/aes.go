package aesx

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"github.com/hootuu/gelato/errors"
	"io"
)

func ToString(src []byte) string {
	return hex.EncodeToString(src)
}

func Encrypt(src []byte, priKey []byte) ([]byte, *errors.Error) {
	if len(src) == 0 {
		return nil, errors.E("src is empty")
	}
	block, err := aes.NewCipher(priKey)
	if err != nil {
		return nil, errors.System("aes.NewCipher([]byte(gAesKey)): " + err.Error())
	}
	padding := aes.BlockSize - len(src)%aes.BlockSize
	paddedPlaintext := append(src, bytes.Repeat([]byte{byte(padding)}, padding)...)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, errors.System("init iv failed: " + err.Error())
	}
	ciphertext := make([]byte, aes.BlockSize+len(paddedPlaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlaintext)
	copy(ciphertext[:aes.BlockSize], iv)
	return ciphertext, nil
}

func Decrypt(src []byte, priKey []byte) ([]byte, *errors.Error) {
	if len(src) < aes.BlockSize {
		return nil, errors.E("src is too short")
	}
	block, err := aes.NewCipher(priKey)
	if err != nil {
		return nil, errors.System("aes.NewCipher([]byte(gAesKey)): " + err.Error())
	}
	iv := src[:aes.BlockSize]
	decrypted := make([]byte, len(src)-aes.BlockSize)
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(decrypted, src[aes.BlockSize:])
	padding := int(decrypted[len(decrypted)-1])
	decrypted = decrypted[:len(decrypted)-padding]
	return decrypted, nil
}
