package crypt

import (
	"encoding/base64"
)

// Base64Encoding 加密
func Base64Encoding(str string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(str))
	return encoded
}

// Base64Decoding 解密
func Base64Decoding(str string) string {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return str
	}
	return string(decoded)
}

// Base64DisplaceEncoding 移位加密
func Base64DisplaceEncoding(str string, right uint8) string {
	bytes := []byte(str)
	for i := range bytes {
		x := bytes[i]
		u := x ^ right
		bytes[i] = u
	}
	encoded := base64.StdEncoding.EncodeToString(bytes)
	return encoded
}

// Base64DisplaceDecoding 移位解密
func Base64DisplaceDecoding(str string, right uint8) string {
	bytes, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return str
	}
	for i := range bytes {
		x := bytes[i]
		u := x ^ right
		bytes[i] = u
	}
	return string(bytes)
}
