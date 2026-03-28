package main

import (
    "strings"
)

func WordsString(s string) []string {
    if s == "" {
        return []string{}
    }
    
    // Replace commas with spaces
    s = strings.ReplaceAll(s, ",", " ")
    
    // Split by whitespace using Fields (handles multiple spaces and removes empties)
    return strings.Fields(s)
}