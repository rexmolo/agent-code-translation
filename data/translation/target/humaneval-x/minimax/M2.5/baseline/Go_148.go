package main

import "fmt"

func Bf(planet1, planet2 string) []string {
	planetNames := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

	// Check if planets are valid
	planet1Valid := false
	planet2Valid := false
	var planet1Index, planet2Index int

	for i, p := range planetNames {
		if p == planet1 {
			planet1Valid = true
			planet1Index = i
		}
		if p == planet2 {
			planet2Valid = true
			planet2Index = i
		}
	}

	if !planet1Valid || !planet2Valid || planet1 == planet2 {
		return []string{}
	}

	if planet1Index < planet2Index {
		return planetNames[planet1Index+1 : planet2Index]
	} else {
		return planetNames[planet2Index+1 : planet1Index]
	}
}

func main() {
	// Test cases
	fmt.Println(Bf("Jupiter", "Neptune")) // [Saturn Uranus]
	fmt.Println(Bf("Earth", "Mercury"))  // [Venus]
	fmt.Println(Bf("Mercury", "Uranus")) // [Venus Earth Mars Jupiter Saturn]
	fmt.Println(Bf("Earth", "Earth"))     // []
	fmt.Println(Bf("Pluto", "Mars"))      // []
}