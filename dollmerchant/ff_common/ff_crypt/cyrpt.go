package ff_crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// SHA1 SHA1哈希加密
func SHA1(plainText []byte) string {
	sha := sha1.New()
	sha.Write(plainText)
	return hex.EncodeToString(sha.Sum(nil))
}

// MD5 MD5哈希加密， 返回32位字符串
func MD5(plainText []byte) string {
	m := md5.New()
	m.Write(plainText)
	return hex.EncodeToString(m.Sum(nil))
}

// AESEncrypt AES加密
func AESEncrypt(plaintext, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	plaintext = pkcsPadding(plaintext, block.BlockSize())
	mode := cipher.NewCBCEncrypter(block, iv)
	crypted := make([]byte, len(plaintext))
	mode.CryptBlocks(crypted, plaintext)
	return crypted, nil
}

// AESDecrypt AES 解密
func AESDecrypt(crypted, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(crypted))
	mode.CryptBlocks(origData, crypted)
	origData = pkcsUnPadding(origData)
	return origData, nil
}

func pkcsPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcsUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
