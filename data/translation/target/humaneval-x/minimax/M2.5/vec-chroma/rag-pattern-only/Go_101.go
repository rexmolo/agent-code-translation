package main

import (
    "strings"
)

func WordsString(s string) []string {
    if s == "" {
        return []string{}
    }

    var result []rune
    for _, r := range s {
        if r == ',' {
            result = append(result, ' ')
        } else {
            result = append(result, r)
        }
    }

    return strings.Fields(string(result))
}