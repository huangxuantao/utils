package encrypt_util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

// 常量 keys iv 用于AES加解密
const keys = "2018201820182018"
const iv = "1234567887654321"

// authKey 用于生成 sys key
const authKey = "VanyrMjsX0av"

// 常量 k 用于Des加密
const k = "lyonsdpy"
var key = []byte(k)

// 加解密填充方法
func pKCS5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func pKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}

func pKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pKCS7UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}

func DesEncrypt(clearText string) (string, error) {
	data := []byte(clearText)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	data = pKCS5Padding(data, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	cipherText := make([]byte, len(data))
	blockMode.CryptBlocks(cipherText, data)
	return hex.EncodeToString(cipherText), nil
}

func DesDecrypt(cipherText string) (string, error) {
	data, _ := hex.DecodeString(cipherText)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	clearText := make([]byte, len(data))
	blockMode.CryptBlocks(clearText, data)
	clearText = pKCS5UnPadding(clearText)
	return string(clearText), nil
}

func GenerateSysKey(sysName string) string {
	h := sha256.New()
	h.Write([]byte(sysName))
	sysKey := h.Sum([]byte(authKey))
	return fmt.Sprintf("%x", sysKey)
}

func AesEncrypt(clearText string) (string, error) {
	data := []byte(clearText)
	block, err := aes.NewCipher([]byte(keys))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	data = pKCS7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(data))
	blockMode.CryptBlocks(cipherText, data)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AesDecrypt(cipherText string) (string, error) {
	data, _ := base64.StdEncoding.DecodeString(cipherText)
	block, err := aes.NewCipher([]byte(keys))
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	clearText := make([]byte, len(data))
	blockMode.CryptBlocks(clearText, data)
	clearText = pKCS7UnPadding(clearText)
	return string(clearText), nil
}

func AesEncryptByKeyIv(clearText string, key string, iv string) (string, error) {
	data := []byte(clearText)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	data = pKCS7Padding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, []byte(iv))
	cipherText := make([]byte, len(data))
	blockMode.CryptBlocks(cipherText, data)
	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func AesDecryptByKeyIv(cipherText string, key string, iv string) (string, error) {
	data, _ := base64.StdEncoding.DecodeString(cipherText)
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	clearText := make([]byte, len(data))
	blockMode.CryptBlocks(clearText, data)
	clearText = pKCS7UnPadding(clearText)
	return string(clearText), nil
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
