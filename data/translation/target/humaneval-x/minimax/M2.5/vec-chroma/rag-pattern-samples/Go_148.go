package main

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

    // Return empty slice if either planet is invalid or if they are the same
    if planet1Index == -1 || planet2Index == -1 || planet1Index == planet2Index {
        return []string{}
    }

    // Return planets between the two planets, sorted by proximity to the sun
    if planet1Index < planet2Index {
        return planetNames[planet1Index+1 : planet2Index]
    } else {
        return planetNames[planet2Index+1 : planet1Index]
    }
}