import (
    "slices"
    "strings"
)

func SortNumbers(numbers string) string {
    valueMap := map[string]int{
        "zero":  0,
        "one":   1,
        "two":   2,
        "three": 3,
        "four":  4,
        "five":  5,
        "six":   6,
        "seven": 7,
        "eight": 8,
        "nine":  9,
    }

    // Split the string and filter empty elements (Python: [x for x in numbers.split(' ') if x])
    var words []string
    for _, x := range strings.Split(numbers, " ") {
        if x != "" {
            words = append(words, x)
        }
    }

    // Sort by numeric value (Python: sorted(..., key=lambda x: value_map[x]))
    slices.SortFunc(words, func(a, b string) int {
        return valueMap[a] - valueMap[b]
    })

    return strings.Join(words, " ")
}