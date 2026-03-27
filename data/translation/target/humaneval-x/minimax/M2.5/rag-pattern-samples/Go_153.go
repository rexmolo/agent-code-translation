package main

import (
    "fmt"
    "unicode"
)

func StrongestExtension(class_name string, extensions []string) string {
    if len(extensions) == 0 {
        return class_name + "."
    }
    
    strong := extensions[0]
    myVal := calculateStrength(extensions[0])
    
    for _, s := range extensions {
        val := calculateStrength(s)
        if val > myVal {
            strong = s
            myVal = val
        }
    }
    
    return class_name + "." + strong
}

func calculateStrength(ext string) int {
    upper := 0
    lower := 0
    for _, r := range ext {
        if unicode.IsUpper(r) {
            upper++
        }
        if unicode.IsLower(r) {
            lower++
        }
    }
    return upper - lower
}

func main() {
    result := StrongestExtension("Slices", []string{"SErviNGSliCes", "Cheese", "StuFfed"})
    fmt.Println(result)
    
    result2 := StrongestExtension("my_class", []string{"AA", "Be", "CC"})
    fmt.Println(result2)
}
