package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func encodeCyclic(s string) string {
	// split string to groups. Each of length 3.
	var groups []string
	for i := 0; i < len(s); i += 3 {
		end := i + 3
		if end > len(s) {
			end = len(s)
		}
		groups = append(groups, s[i:end])
	}

	// cycle elements in each group. Unless group has fewer elements than 3.
	for i, group := range groups {
		if len(group) == 3 {
			groups[i] = group[1:] + group[0:1]
		}
	}
	return strings.Join(groups, "")
}

func DecodeCyclic(s string) string {
	// Decoding is done by applying encode twice
	return encodeCyclic(encodeCyclic(s))
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	result := DecodeCyclic(input)
	fmt.Println(result)
}
