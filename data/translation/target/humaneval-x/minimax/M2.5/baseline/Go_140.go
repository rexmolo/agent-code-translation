package main

import (
    "strings"
)

func FixSpaces(text string) string {
    var builder strings.Builder
    i := 0
    start := 0
    end := 0

    for i < len(text) {
        if text[i] == ' ' {
            end++
        } else {
            spaceCount := end - start
            if spaceCount > 2 {
                builder.WriteByte('-')
            } else if spaceCount > 0 {
                for j := 0; j < spaceCount; j++ {
                    builder.WriteByte('_')
                }
            }
            builder.WriteByte(text[i])
            start = i + 1
            end = i + 1
        }
        i++
    }

    spaceCount := end - start
    if spaceCount > 2 {
        builder.WriteByte('-')
    } else if spaceCount > 0 {
        for j := 0; j < spaceCount; j++ {
            builder.WriteByte('_')
        }
    }

    return builder.String()
}