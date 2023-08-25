package crypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

// AesEncrypt  通用加密函数
func AesEncrypt(data []byte, aesKey string) []byte {
	return aesEncryptCBC(data, []byte(aesKey))
}

// AesDecrypt  通用解密函数
func AesDecrypt(data []byte, aesKey string) ([]byte, error) {
	ret, err := aesDecryptCBC(data, []byte(aesKey))
	return ret, err
}

func aesEncryptCBC(origData []byte, key []byte) (encrypted []byte) {
	block, _ := aes.NewCipher(key)
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return encrypted
}

func aesDecryptCBC(encrypted []byte, key []byte) (decrypted []byte, err error) {
	block, _ := aes.NewCipher(key)                              // 分组秘钥
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted, err = pkcs5UnPadding(decrypted)                  // 去除补全码
	return decrypted, err
}

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) ([]byte, error) {
	length := len(origData)
	if length == 0 {
		return nil, errors.New("decrypt error")
	}
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)], nil
}
