package main

import (
	"fmt"
)

// NumericalLetterGrade converts a slice of GPA scores into a slice of letter grades.
// The grading scale is based on a specific algorithm.
func NumericalLetterGrade(grades []float64) []string {
	letterGrade := make([]string, 0, len(grades))
	for _, gpa := range grades {
		if gpa == 4.0 {
			letterGrade = append(letterGrade, "A+")
		} else if gpa > 3.7 {
			letterGrade = append(letterGrade, "A")
		} else if gpa > 3.3 {
			letterGrade = append(letterGrade, "A-")
		} else if gpa > 3.0 {
			letterGrade = append(letterGrade, "B+")
		} else if gpa > 2.7 {
			letterGrade = append(letterGrade, "B")
		} else if gpa > 2.3 {
			letterGrade = append(letterGrade, "B-")
		} else if gpa > 2.0 {
			letterGrade = append(letterGrade, "C+")
		} else if gpa > 1.7 {
			letterGrade = append(letterGrade, "C")
		} else if gpa > 1.3 {
			letterGrade = append(letterGrade, "C-")
		} else if gpa > 1.0 {
			letterGrade = append(letterGrade, "D+")
		} else if gpa > 0.7 {
			letterGrade = append(letterGrade, "D")
		} else if gpa > 0.0 {
			letterGrade = append(letterGrade, "D-")
		} else {
			letterGrade = append(letterGrade, "E")
		}
	}
	return letterGrade
}

func main() {
	// Example from the original problem description.
	exampleGrades := []float64{4.0, 3, 1.7, 2, 3.5}
	result := NumericalLetterGrade(exampleGrades)
	// The Go code uses %#v to print a Go-syntax representation of the value.
	fmt.Printf("Input: %v\n", exampleGrades)
	fmt.Printf("Result: %v\n", result)
}
