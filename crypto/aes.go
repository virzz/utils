package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// Padding PKCS5
func Padding(src []byte, blockSize int) []byte {
	p := blockSize - len(src)%blockSize
	return append(src, bytes.Repeat([]byte{byte(p)}, p)...)
}

// UnPadding PKCS5
func UnPadding(src []byte) []byte {
	l := len(src)
	if n := int(src[l-1]); n <= l {
		return src[:l-n]
	}
	return src
}

// 加密
func AesEncrypt(data, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	iv := make([]byte, aes.BlockSize)
	io.ReadFull(rand.Reader, iv)
	data = Padding(data, aes.BlockSize)
	buf := make([]byte, len(data))
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(buf, data)
	return buf, iv, nil
}

// 解密
func AesDecrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// data = Padding(data, aes.BlockSize)
	buf := make([]byte, len(data))
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(buf, data)
	return UnPadding(buf), nil
}
