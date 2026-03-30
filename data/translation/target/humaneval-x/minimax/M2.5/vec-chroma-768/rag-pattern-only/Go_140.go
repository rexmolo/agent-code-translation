package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FixSpaces(text string) string {
	var result strings.Builder
	i := 0
	start := 0
	end := 0

	for i < len(text) {
		if text[i] == ' ' {
			end++
		} else {
			if end-start > 2 {
				result.WriteString("-")
				result.WriteByte(text[i])
			} else if end-start > 0 {
				for j := 0; j < end-start; j++ {
					result.WriteByte('_')
				}
				result.WriteByte(text[i])
			} else {
				result.WriteByte(text[i])
			}
			start = i + 1
			end = i + 1
		}
		i++
	}

	if end-start > 2 {
		result.WriteString("-")
	} else if end-start > 0 {
		result.WriteString("_")
	}

	return result.String()
}


func main() {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	fmt.Println(FixSpaces(text))
}
