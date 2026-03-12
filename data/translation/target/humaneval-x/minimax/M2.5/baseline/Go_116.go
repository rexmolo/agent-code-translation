package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func countOnes(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

type byOnesAndValue []int

func (a byOnesAndValue) Len() int      { return len(a) }
func (a byOnesAndValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a byOnesAndValue) Less(i, j int) bool {
	iOnes := countOnes(a[i])
	jOnes := countOnes(a[j])
	if iOnes != jOnes {
		return iOnes < jOnes
	}
	return a[i] < a[j]
}

func SortArray(arr []int) []int {
	sortedArr := make([]int, len(arr))
	copy(sortedArr, arr)
	sort.Sort(byOnesAndValue(sortedArr))
	return sortedArr
}

func main() {
	// Read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter array elements separated by spaces (e.g., 1 5 2 3 4): ")
	scanner.Scan()
	input := scanner.Text()
	
	// Parse input
	strs := strings.Fields(input)
	arr := make([]int, 0, len(strs))
	for _, s := range strs {
		n, err := strconv.Atoi(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing number: %v\n", err)
			return
		}
		arr = append(arr, n)
	}
	
	result := SortArray(arr)
	fmt.Printf("Sorted array: %v\n", result)
}
