package main

import (
	"fmt"
	"strings"
)

// ParseMusic parses a string representing musical notes in a special ASCII format.
// It returns a slice of integers corresponding to the beat duration of each note.
// Legend:
// 'o'  - whole note, lasts four beats
// 'o|' - half note, lasts two beats
// '.|' - quater note, lasts one beat
func ParseMusic(musicString string) []int {
	noteMap := map[string]int{
		"o":  4,
		"o|": 2,
		".|": 1,
	}

	// Split the input string by spaces to get individual note representations.
	notes := strings.Split(musicString, " ")

	// Initialize a slice to hold the beat counts.
	// We can pre-allocate with a capacity for slight optimization.
	beats := make([]int, 0, len(notes))

	for _, note := range notes {
		// The python version `if x` in the list comprehension filters out empty strings,
		// which can result from multiple spaces. This check does the same.
		if note == "" {
			continue
		}

		// Look up the note in the map and append its beat value.
		// The original python code assumes all notes are valid keys.
		// This Go code does the same; accessing a non-existent key would append 0.
		beats = append(beats, noteMap[note])
	}

	return beats
}

func main() {
	testString := "o o| .| o| o| .| .| .| .| o o"
	result := ParseMusic(testString)
	fmt.Println(result)

	// Another test case with extra spaces
	testString2 := "o  o| .|  o o "
	result2 := ParseMusic(testString2)
	fmt.Println(result2)
}