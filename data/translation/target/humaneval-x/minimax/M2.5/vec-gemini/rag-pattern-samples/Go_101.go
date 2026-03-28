package main

import (
    "strings"
)

func WordsString(s string) []string {
    if s == "" {
        return []string{}
    }

    // Replace commas with spaces, then split on whitespace
    s = strings.Replace(s, ",", " ", -1)
    return strings.Fields(s)
}