package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
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
	fmt.Println(StringToMd5("Hello world"))
	fmt.Println(StringToMd5(""))
}