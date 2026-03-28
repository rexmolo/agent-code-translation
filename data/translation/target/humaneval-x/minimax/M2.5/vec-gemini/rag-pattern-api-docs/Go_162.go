package main

import (
    "crypto/md5"
    "encoding/hex"
)

func StringToMd5(text string) interface{} {
    if text == "" {
        return nil
    }
    hasher := md5.New()
    hasher.Write([]byte(text))
    return hex.EncodeToString(hasher.Sum(nil))
}