package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"github.com/pkg/errors"
)

// 强行使用 AES-256-CBC
const BlockSize = 32

// Padding PKCS5
func Padding(src []byte, blockSize int) []byte {
	p := blockSize - len(src)%blockSize
	return append(src, bytes.Repeat([]byte{byte(p)}, p)...)
}

// UnPadding PKCS5
func UnPadding(src []byte) []byte {
	return src[:len(src)-int(src[len(src)-1])]
}

// 加密
func AesEncrypt(data, key []byte) ([]byte, []byte, error) {
	if len(key) != BlockSize {
		return nil, nil, errors.Errorf("crypto/aes: invalid key size %d", BlockSize)
	}
	block, _ := aes.NewCipher(key)
	iv := make([]byte, BlockSize)
	io.ReadFull(rand.Reader, iv)
	data = Padding(data, block.BlockSize())
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(data, data)
	return data, iv, nil
}

// 解密
func AesDecrypt(data, key, iv []byte) ([]byte, error) {
	if len(key) != BlockSize {
		return nil, errors.Errorf("crypto/aes: invalid key size %d", BlockSize)
	}
	if len(iv) != BlockSize {
		return nil, errors.Errorf("cipher.NewCBCDecrypter: IV length must equal block size %d", BlockSize)
	}
	block, _ := aes.NewCipher(key)
	cipher.NewCBCDecrypter(block, iv).CryptBlocks(data, data)
	return UnPadding(data), nil
}
