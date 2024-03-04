package crypto

import (
	"encoding/pem"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/wenzhenxi/gorsa"
)

// 私钥加密
func RsaPemPriEncrypt(src string, privateKey string) (string, error) {
	var srcArr []string
	for {
		if len(src) > 100 {
			srcArr = append(srcArr, src[:100])
			src = src[100:]
		} else {
			srcArr = append(srcArr, src)
			break
		}
	}

	var encArr []string
	for _, seg := range srcArr {
		enc, err := gorsa.PriKeyEncrypt(seg, privateKey)
		if err != nil {
			return "", err
		}
		encArr = append(encArr, enc)
	}
	return strings.Join(encArr, "\n"), nil
}

// 公钥解密
func RsaPemPubDecrypt(enc string, publicKey string) (string, error) {
	var srcArr []string
	encArr := strings.Split(enc, "\n")
	for _, seq := range encArr {
		src, err := gorsa.PublicDecrypt(seq, publicKey)
		if err != nil {
			return "", err
		}
		srcArr = append(srcArr, src)
	}
	return strings.Join(srcArr, ""), nil
}

// 公钥加密
func RsaPemPubEncrypt(src string, publicKey string) (string, error) {
	var srcArr []string
	for {
		if len(src) > 100 {
			srcArr = append(srcArr, src[:100])
			src = src[100:]
		} else {
			srcArr = append(srcArr, src)
			break
		}
	}

	var encArr []string
	for _, seg := range srcArr {
		enc, err := gorsa.PublicEncrypt(seg, publicKey)
		if err != nil {
			return "", err
		}
		encArr = append(encArr, enc)
	}
	return strings.Join(encArr, "\n"), nil
}

// 私钥解密
func RsaPemPriDecrypt(enc string, privateKey string) (string, error) {
	var srcArr []string
	encArr := strings.Split(enc, "\n")
	for _, seq := range encArr {
		src, err := gorsa.PriKeyDecrypt(seq, privateKey)
		if err != nil {
			return "", err
		}
		srcArr = append(srcArr, src)
	}
	return strings.Join(srcArr, ""), nil
}

func ReadPublicKey(filename string) ([]byte, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return PublicBytes(buf)
}

func PublicBytes(buf []byte) ([]byte, error) {
	p, _ := pem.Decode(buf)
	if p.Type != "PUBLIC KEY" {
		return nil, errors.Errorf("invalid public key type: %s", p.Type)
	}
	return p.Bytes, nil
}
