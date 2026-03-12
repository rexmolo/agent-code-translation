package main

import (
	"crypto/md5"
	"encoding/hex"
)

func StringToMd5(text string) interface{} {
	if text == "" {
		return nil
	}
	data := []byte(text)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}
