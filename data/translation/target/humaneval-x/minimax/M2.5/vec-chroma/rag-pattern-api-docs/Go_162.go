package main

import (
    "crypto/md5"
    "encoding/hex"
)

// StringToMd5 computes the MD5 hash of a string and returns it as a hex string.
// Returns nil if the input string is empty.
func StringToMd5(text string) interface{} {
    if text == "" {
        return nil
    }
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}

func main() {
    // Example usage
    result := StringToMd5("Hello world")
    println(result)
    
    // Test empty string
    println(StringToMd5(""))
}