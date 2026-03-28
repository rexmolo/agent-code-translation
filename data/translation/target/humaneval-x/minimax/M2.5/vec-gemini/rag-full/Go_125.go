package main

import (
    "strings"
)

func SplitWords(txt string) interface{} {
    if strings.Contains(txt, " ") {
        return strings.Fields(txt)
    } else if strings.Contains(txt, ",") {
        return strings.Fields(strings.ReplaceAll(txt, ",", " "))
    } else {
        count := 0
        for _, ch := range txt {
            if ch >= 'a' && ch <= 'z' && (int(ch)-int('a'))%2 == 0 {
                count++
            }
        }
        return count
    }
}