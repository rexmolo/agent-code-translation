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
	data := []byte(text)
	hash := md5.Sum(data)
	return hex.EncodeToString(hash[:])
}

func main() {
	var input string
	fmt.Scanln(&input)
	result := StringToMd5(input)
	if result != nil {
		fmt.Println(result)
	}
}