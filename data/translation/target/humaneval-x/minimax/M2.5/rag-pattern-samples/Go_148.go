package main

import (
    "fmt"
)

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

    // Return empty slice if either planet is invalid or they are the same
    if planet1Index == -1 || planet2Index == -1 || planet1 == planet2 {
        return []string{}
    }

    // Return planets between the two (exclusive), sorted by proximity to the Sun
    var start, end int
    if planet1Index < planet2Index {
        start = planet1Index + 1
        end = planet2Index
    } else {
        start = planet2Index + 1
        end = planet1Index
    }

    return planetNames[start:end]
}

func main() {
    // Test cases
    fmt.Println(Bf("Jupiter", "Neptune")) 
    fmt.Println(Bf("Earth", "Mercury"))   
    fmt.Println(Bf("Mercury", "Uranus"))
}