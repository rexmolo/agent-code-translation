package main

import "fmt"

func NumericalLetterGrade(grades []float64) []string {
	letterGrades := make([]string, 0, len(grades))

	for _, gpa := range grades {
		if gpa == 4.0 {
			letterGrades = append(letterGrades, "A+")
		} else if gpa > 3.7 {
			letterGrades = append(letterGrades, "A")
		} else if gpa > 3.3 {
			letterGrades = append(letterGrades, "A-")
		} else if gpa > 3.0 {
			letterGrades = append(letterGrades, "B+")
		} else if gpa > 2.7 {
			letterGrades = append(letterGrades, "B")
		} else if gpa > 2.3 {
			letterGrades = append(letterGrades, "B-")
		} else if gpa > 2.0 {
			letterGrades = append(letterGrades, "C+")
		} else if gpa > 1.7 {
			letterGrades = append(letterGrades, "C")
		} else if gpa > 1.3 {
			letterGrades = append(letterGrades, "C-")
		} else if gpa > 1.0 {
			letterGrades = append(letterGrades, "D+")
		} else if gpa > 0.7 {
			letterGrades = append(letterGrades, "D")
		} else if gpa > 0.0 {
			letterGrades = append(letterGrades, "D-")
		} else {
			letterGrades = append(letterGrades, "E")
		}
	}

	return letterGrades
}

func main() {
	// Test the function with example input
	grades := []float64{4.0, 3.0, 1.7, 2.0, 3.5}
	result := NumericalLetterGrade(grades)
	fmt.Println(result)
}