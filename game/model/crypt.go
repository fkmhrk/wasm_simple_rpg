package model

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"github.com/pkg/errors"
)

const (
	encKey = "1234567890123456789012345678901234567890123456789012345678901234"
)

func Encrypt(value []byte) ([]byte, string, error) {
	key, _ := hex.DecodeString(encKey)
	iv, _ := hex.DecodeString("1234567890ABCDEF1234567890ABCDEF")
	if len(iv) != aes.BlockSize {
		return nil, "", errors.Errorf("illegal initial vector size [%d]byte. initial vector size must be [%d]byte", len(iv), aes.BlockSize)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, "", errors.Wrap(err, "failed to create AES cipher block")
	}
	encrypter := cipher.NewCBCEncrypter(block, iv)
	padded := addPadding(value)
	encrypted := make([]byte, len(padded))
	encrypter.CryptBlocks(encrypted, padded)
	return encrypted, "1234567890ABCDEF1234567890ABCDEF", nil
}

func Decrypt(data []byte, ivStr string) ([]byte, error) {
	key, _ := hex.DecodeString(encKey)
	iv, _ := hex.DecodeString(ivStr)
	if len(iv) != aes.BlockSize {
		return nil, errors.Errorf("illegal initial vector size [%d]byte. initial vector size must be [%d]byte", len(iv), aes.BlockSize)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create AES cipher block")
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	plain := make([]byte, len(data))
	mode.CryptBlocks(plain, data)
	return removePadding(plain), nil
}

func addPadding(value []byte) []byte {
	padSize := aes.BlockSize - (len(value) % aes.BlockSize)
	pad := bytes.Repeat([]byte{byte(padSize)}, padSize)
	return append(value, pad...)
}

func removePadding(value []byte) []byte {
	padSize := int(value[len(value)-1])
	return value[:len(value)-padSize]
}
