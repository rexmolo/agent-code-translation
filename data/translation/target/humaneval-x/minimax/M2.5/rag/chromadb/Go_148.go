package main

import (
    "fmt"
    "slices"
)

func Bf(planet1, planet2 string) []string {
    planetNames := []string{"Mercury", "Venus", "Earth", "Mars", "Jupiter", "Saturn", "Uranus", "Neptune"}

    // Check if either planet is not in the list or if they are the same
    if !slices.Contains(planetNames, planet1) || !slices.Contains(planetNames, planet2) || planet1 == planet2 {
        return []string{}
    }

    planet1Index := slices.Index(planetNames, planet1)
    planet2Index := slices.Index(planetNames, planet2)

    if planet1Index < planet2Index {
        return planetNames[planet1Index+1 : planet2Index]
    } else {
        return planetNames[planet2Index+1 : planet1Index]
    }
}

func main() {
    // Test cases
    fmt.Println(Bf("Jupiter", "Neptune"))  // [Saturn Uranus]
    fmt.Println(Bf("Earth", "Mercury"))    // [Venus]
    fmt.Println(Bf("Mercury", "Uranus"))   // [Venus Earth Mars Jupiter Saturn]
}