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

func main() {
	// Test cases
	result1 := StringToMd5("Hello world")
	println(result1)

	result2 := StringToMd5("")
	println(result2)
}
