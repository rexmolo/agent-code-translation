package main

import (
	"crypto/md5"
	"encoding/hex"
)

func StringToMd5(text string) interface{} {
	if text == "" {
		return nil
	}
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}