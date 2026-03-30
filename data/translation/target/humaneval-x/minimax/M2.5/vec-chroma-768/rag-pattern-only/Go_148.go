package main

import "fmt"

func Bf(planet1, planet2 string) []string {
	planetNames := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// Find indices of both planets
	planet1Index := -1
	planet2Index := -1

	for i, p := range planetNames {
		if p == planet1 {
			planet1Index = i
		}
		if p == planet2 {
			planet2Index = i
		}
	}

	// Return empty slice if either planet is invalid or if they're the same
	if planet1Index == -1 || planet2Index == -1 || planet1 == planet2 {
		return []string{}
	}

	// Return planets between the two (exclusive of both)
	if planet1Index < planet2Index {
		return planetNames[planet1Index+1 : planet2Index]
	} else {
		return planetNames[planet2Index+1 : planet1Index]
	}
}

func main() {
	fmt.Println(Bf("Jupiter", "Neptune"))
	fmt.Println(Bf("Earth", "Mercury"))
	fmt.Println(Bf("Mercury", "Uranus"))
}