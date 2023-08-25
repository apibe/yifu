package crypt

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(bytes []byte) string {
	h := md5.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
