package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// IsBored counts the number of "boredoms". A boredom is defined as a sentence
// fragment that begins with the characters "I ". Sentences are delimited by '.',
// '?' or '!'.
func IsBored(S string) int {
	// The Python code uses re.split(r'[.?!]\s*', S).
	// This splits the string by '.', '?', or '!' followed by optional whitespace.
	re := regexp.MustCompile(`[.?!]\s*`)

	// Split the input string into sentences. The -1 argument means to find all matches.
	// This behavior is equivalent to Python's re.split.
	sentences := re.Split(S, -1)

	count := 0
	// The Python code `sum(sentence[0:2] == 'I ' for sentence in sentences)`
	// checks if the slice of the first two characters is exactly "I ".
	// It does not trim leading whitespace from each sentence fragment.
	for _, sentence := range sentences {
		// strings.HasPrefix is the idiomatic and safe Go equivalent for checking
		// if a string starts with a specific prefix. It correctly handles
		// strings shorter than the prefix, matching the Python slice behavior.
		if strings.HasPrefix(sentence, "I ") {
			count++
		}
	}

	return count
}

// main function to make the code runnable. It reads a line from stdin,
// processes it with IsBored, and prints the result to stdout.
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line := scanner.Text()
		result := IsBored(line)
		fmt.Println(result)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading from stdin:", err)
	}
}
