package main

import (
	"fmt"
)

func NumericalLetterGrade(grades []float64) []string {
	letterGrades := make([]string, len(grades))

	for i, gpa := range grades {
		if gpa == 4.0 {
			letterGrades[i] = "A+"
		} else if gpa > 3.7 {
			letterGrades[i] = "A"
		} else if gpa > 3.3 {
			letterGrades[i] = "A-"
		} else if gpa > 3.0 {
			letterGrades[i] = "B+"
		} else if gpa > 2.7 {
			letterGrades[i] = "B"
		} else if gpa > 2.3 {
			letterGrades[i] = "B-"
		} else if gpa > 2.0 {
			letterGrades[i] = "C+"
		} else if gpa > 1.7 {
			letterGrades[i] = "C"
		} else if gpa > 1.3 {
			letterGrades[i] = "C-"
		} else if gpa > 1.0 {
			letterGrades[i] = "D+"
		} else if gpa > 0.7 {
			letterGrades[i] = "D"
		} else if gpa > 0.0 {
			letterGrades[i] = "D-"
		} else {
			letterGrades[i] = "E"
		}
	}

	return letterGrades
}

func main() {
	// Example test
	grades := []float64{4.0, 3.0, 1.7, 2.0, 3.5}
	result := NumericalLetterGrade(grades)
	fmt.Println(result)
}